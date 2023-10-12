package rest

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/middleware"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/user"
	"github.com/irdaislakhuafa/go-argon2/argon2"
	"github.com/labstack/echo/v4"
)

type UserInterface interface {
	GetList(ctx context.Context) func(c echo.Context) error
	Update(ctx context.Context) func(c echo.Context) error
}

type restUser struct {
	rest *Rest
}

func NewUser(rest *Rest, ctx context.Context) UserInterface {
	user := &restUser{
		rest: rest,
	}

	rest.echo.GET("/api/users", user.GetList(ctx), middleware.JWT(rest.cfg, middleware.JWTMiddlewareOption{RoleNames: []string{"admin"}}))
	return user
}

func (u *restUser) GetList(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[[]*generated.User]{}
		body := new(struct {
			IsDeleted bool `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		listUser, err := u.rest.client.User.Query().Where(user.IsDeleted(body.IsDeleted)).All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = listUser
		return c.JSON(http.StatusOK, result)
	}
}

func (u *restUser) Update(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[*generated.User]{}
		body := new(struct {
			ID       string `validate:"required"`
			Name     string `validate:"required"`
			Email    string `validate:"required,email"`
			Password string `validate:"required"`
		})
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		hashedPassword, err := argon2.HashArgon2([]byte(body.Password))
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		tx, err := u.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}
		defer tx.Rollback()

		user, err := tx.User.
			UpdateOneID(body.ID).
			SetName(body.Name).
			SetEmail(body.Email).
			SetPassword(hashedPassword).
			SetUpdatedAt(time.Now()).
			SetUpdatedBy(c.Get("user_id").(string)).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		result.Data = user
		return c.JSON(http.StatusOK, result)
	}
}
