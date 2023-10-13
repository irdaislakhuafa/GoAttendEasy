package rest

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/middleware"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/division"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/employee"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/user"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/customvalidator"
	"github.com/labstack/echo/v4"
)

type EmployeeInterface interface {
	Create(ctx context.Context) func(c echo.Context) error
	GetList(ctx context.Context) func(c echo.Context) error
	Get(ctx context.Context) func(c echo.Context) error
	Update(ctx context.Context) func(c echo.Context) error
	Delete(ctx context.Context) func(c echo.Context) error
}
type restEmployee struct {
	rest *Rest
}

func NewEmployee(rest *Rest, ctx context.Context) EmployeeInterface {
	employee := &restEmployee{
		rest: rest,
	}

	rest.echo.POST("/api/employees", employee.Create(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.GET("/api/employees", employee.GetList(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.GET("/api/employees/:id", employee.Get(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{}))
	rest.echo.PUT("/api/employees", employee.Update(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	rest.echo.DELETE("/api/employees", employee.Delete(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	return employee
}

func (e *restEmployee) Create(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			UserID     string `validate:"required"`
			DivisionID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := e.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		if _, err := tx.Division.Query().Where(division.ID(body.DivisionID)).First(ctx); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if _, err := tx.User.Query().Where(user.ID(body.UserID)).First(ctx); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		employee, err := tx.Employee.Create().
			SetID(uuid.NewString()).
			SetUserID(body.UserID).
			SetDivisionID(body.DivisionID).
			SetCreatedAt(time.Now()).
			SetCreatedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) GetList(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[[]*generated.Employee]{}
		body := new(struct {
			IsDeleted bool
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		listEmployee, err := e.rest.client.Employee.Query().Where(employee.IsDeleted(body.IsDeleted)).All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = listEmployee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) Get(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		id := c.Param("id")

		employee, err := e.rest.client.Employee.Get(ctx, id)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) Update(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			ID         string `validate:"required"`
			UserID     string `validate:"required"`
			DivisionID string `validate:"required"`
		})

		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := e.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		if _, err := tx.User.Get(ctx, body.UserID); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		employee, err := tx.Employee.UpdateOneID(body.ID).
			SetUserID(body.UserID).
			SetDivisionID(body.DivisionID).
			SetUpdatedAt(time.Now()).
			SetUpdatedBy(c.Get("user_id").(string)).
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

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}

func (e *restEmployee) Delete(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.Employee]{}
		body := new(struct {
			ID string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := e.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		employee, err := tx.Employee.UpdateOneID(body.ID).
			SetDeletedAt(time.Now()).
			SetDeletedBy(c.Get("user_id").(string)).
			SetIsDeleted(true).
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

		result.Data = employee
		return c.JSON(http.StatusOK, result)
	}
}
