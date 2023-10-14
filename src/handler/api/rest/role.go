package rest

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/GoAttendEasy/src/entity"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/middleware"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/role"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/customvalidator"
	"github.com/labstack/echo/v4"
)

type RoleInterface interface {
	Create(ctx context.Context) func(c echo.Context) error
	GetList(ctx context.Context) func(c echo.Context) error
	Get(ctx context.Context) func(c echo.Context) error
	Update(ctx context.Context) func(c echo.Context) error
	Delete(ctx context.Context) func(c echo.Context) error
}
type restRole struct {
	rest *Rest
}

func NewRole(rest *Rest, ctx context.Context) RoleInterface {
	role := &restRole{
		rest: rest,
	}

	rest.echo.POST("/api/roles", role.Create(ctx))
	rest.echo.GET("/api/roles", role.GetList(ctx))
	rest.echo.GET("/api/roles/:id", role.Get(ctx))
	rest.echo.PUT("/api/roles", role.Update(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.DELETE("/api/roles", role.Delete(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	return role
}

func (r *restRole) Create(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		body := new(struct {
			Name        string `validate:"required"`
			Description string `validate:"required"`
		})
		result := response.ResponseData[*generated.Role]{}
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, customvalidator.GetErrorsMessage(err)...)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := r.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		role, err := tx.Role.Create().
			SetID(uuid.NewString()).
			SetName(body.Name).
			SetDescription(body.Description).
			SetCreatedBy(r.rest.cfg.App.DefaultCreation).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = role
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restRole) GetList(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[[]*generated.Role]{}
		body := new(struct {
			IsDeleted bool
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		roles, err := r.rest.client.Role.Query().Where(role.IsDeleted(body.IsDeleted)).All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = roles
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restRole) Update(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Role]{}
		body := new(struct {
			ID          string `validate:"required"`
			Name        string `validate:"required"`
			Description string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, customvalidator.GetErrorsMessage(err)...)
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		claims, isOk := c.Get("claims").(*entity.Claims)
		if !isOk {
			result.Error = append(result.Error, map[string]string{"message": "claims body is not valid"})
			return c.JSON(http.StatusUnauthorized, result)
		}

		tx, err := r.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": "claims body is not valid"})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		role, err := tx.Role.UpdateOneID(body.ID).
			SetName(body.Name).
			SetDescription(body.Description).
			SetUpdatedAt(time.Now()).
			SetUpdatedBy(claims.UserID).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = role
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restRole) Delete(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Role]{}
		body := new(struct {
			ID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := r.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		claims, isOk := c.Get("claims").(*entity.Claims)
		if !isOk {
			return c.JSON(http.StatusBadRequest, "claims body is not valid")
		}

		role, err := tx.Role.UpdateOneID(body.ID).
			SetIsDeleted(true).
			SetDeletedAt(time.Now()).
			SetDeletedBy(claims.UserID).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = role
		return c.JSON(http.StatusOK, result)
	}
}

func (r *restRole) Get(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Role]{}
		id := c.Param("id")

		role, err := r.rest.client.Role.Get(ctx, id)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = role
		return c.JSON(http.StatusOK, result)
	}
}
