package rest

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-playground/validator/v10"
	"github.com/irdaislakhuafa/GoAttendEasy/src/entity"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/request"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/role"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/user"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/customvalidator"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/tokens"
	"github.com/irdaislakhuafa/go-argon2/argon2"
	"github.com/labstack/echo/v4"
)

type AuthInterface interface {
	Register(ctx context.Context) func(c echo.Context) error
	Login(ctx context.Context) func(c echo.Context) error
}

type restAuth struct {
	rest *Rest
}

func NewAuth(rest *Rest, ctx context.Context) AuthInterface {
	auth := &restAuth{
		rest: rest,
	}
	rest.echo.POST("/api/auth/register", auth.Register(ctx))
	rest.echo.POST("/api/auth/login", auth.Login(ctx))
	return auth
}

func (a *restAuth) Register(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		// parse json body
		result := response.ResponseData[generated.User]{}
		body := new(request.AuthRegister)
		if err := c.Bind(body); err != nil {
			return err
		}

		// validate
		if err := c.Validate(body); err != nil {
			if err, isOk := err.(validator.ValidationErrors); isOk {
				result := response.ResponseData[any]{
					Error: customvalidator.GetErrorsMessage(err),
				}
				return c.JSON(http.StatusBadRequest, result)
			}
			return err
		}

		// using tx
		tx, err := a.rest.client.BeginTx(ctx, &sql.TxOptions{})
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}
		defer tx.Rollback()

		// find role by list role id
		role, err := tx.Role.Query().Where(role.ID(body.RoleID)).First(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		// hash user password
		hashedPassword, err := argon2.HashArgon2([]byte(body.Password))
		if err != nil {
			return err
		}

		// create user
		user, err := tx.User.Create().
			SetName(body.Name).
			SetEmail(body.Email).
			SetPassword(hashedPassword).
			SetRoleID(role.ID).
			SetCreatedAt(time.Now()).
			SetCreatedBy(a.rest.cfg.App.DefaultCreation).
			Save(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		if err := tx.Commit(); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusInternalServerError, result)
		}

		result = response.ResponseData[generated.User]{Data: *user}
		return c.JSON(http.StatusOK, result)
	}
}

func (a *restAuth) Login(ctx context.Context) func(c echo.Context) error {
	return func(c echo.Context) error {
		result := response.ResponseData[response.AuthToken]{}
		body := new(request.AuthLogin)
		if err := c.Bind(body); err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			return c.JSON(http.StatusBadRequest, result)
		}

		if err := c.Validate(body); err != nil {
			result.Error = customvalidator.GetErrorsMessage(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		user, err := a.rest.client.User.Query().
			Where(user.EmailEqualFold(body.Email)).
			Select(user.FieldPassword).
			Select(user.FieldRoleID).
			First(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		isOk, err := argon2.CompareArgon2(body.Password, user.Password)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		if !isOk {
			result.Error = append(result.Error, map[string]string{"message": fmt.Sprintf("password is not match")})
			return c.JSON(http.StatusBadRequest, result)
		}

		role, err := a.rest.client.Role.
			Query().
			Where(role.ID(user.RoleID)).
			Select(role.FieldName).
			First(ctx)
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		expiredAt := time.Now().Add(time.Minute * time.Duration(a.rest.cfg.App.JWT.ExpiredInMinute))
		claims := &entity.Claims{
			UserID:   user.ID,
			RoleName: role.Name,
			StandardClaims: jwt.StandardClaims{
				IssuedAt:  jwt.Now(),
				ExpiresAt: jwt.At(expiredAt),
			},
		}

		tokenString, err := tokens.NewJWTToken(claims, []byte(a.rest.cfg.App.JWT.Secret))
		if err != nil {
			result.Error = append(result.Error, map[string]string{"message": err.Error()})
			c.Logger().Error(err)
			return c.JSON(http.StatusBadRequest, result)
		}

		token := response.AuthToken{
			Token: tokenString,
		}

		result.Data = token
		result.Message = fmt.Sprintf("token is valid until %v", expiredAt.Format("15:04:05 02/01/2006"))
		return c.JSON(http.StatusOK, result)
	}
}
