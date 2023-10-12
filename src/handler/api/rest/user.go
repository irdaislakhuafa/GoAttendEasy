package rest

import (
	"context"
	"net/http"

	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/middleware"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/labstack/echo/v4"
)

type UserInterface interface {
	GetList(ctx context.Context) func(c echo.Context) error
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
		listUser, err := u.rest.client.User.Query().All(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Data = listUser
		return c.JSON(http.StatusOK, result)
	}
}
