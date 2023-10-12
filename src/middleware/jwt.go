package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/irdaislakhuafa/GoAttendEasy/src/entity"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/model/rest/response"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/operator"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/tokens"
	"github.com/labstack/echo/v4"
)

type JWTMiddlewareOption struct {
	RoleNames []string
}

func JWT(cfg *config.AppConfig, opts JWTMiddlewareOption) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			result := response.ResponseData[any]{}
			authorization := c.Request().Header.Get("authorization")
			if strings.Trim(authorization, " ") == "" {
				result.Error = append(result.Error, map[string]string{"message": "missing jwt token"})
				return c.JSON(http.StatusUnauthorized, result)
			}

			bearer := "bearer "
			if br := strings.ToLower(authorization[:len(bearer)]); br != bearer {
				result.Error = append(result.Error, map[string]string{"string": fmt.Sprintf("authorization header must start with '%v'", bearer)})
				return c.JSON(http.StatusUnauthorized, result)
			}

			tokenString := authorization[len(bearer):]

			token, err := tokens.ValidateJWTToken(tokenString, []byte(cfg.App.JWT.Secret), &entity.Claims{})
			if err != nil {
				result.Error = append(result.Error, map[string]string{"message": err.Error()})
				c.Logger().Errorf(err.Error())
				return c.JSON(http.StatusUnauthorized, result)
			}

			claims, err := tokens.GetClaimsOfJWTToken[*entity.Claims](*token)
			if err != nil {
				result.Error = append(result.Error, map[string]string{"message": err.Error()})
				c.Logger().Errorf(err.Error())
				return c.JSON(http.StatusUnauthorized, result)
			}

			isAllowed := operator.Ternary(len(opts.RoleNames) >= 1, false, true)
			for _, roleName := range opts.RoleNames {
				if strings.EqualFold(roleName, claims.RoleName) {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				result.Error = append(result.Error, map[string]string{"message": fmt.Sprintf("you are not permited to access this resource, allowed role is %s", strings.Join(opts.RoleNames, "/"))})
				return c.JSON(http.StatusUnauthorized, result)
			}

			c.Set("claims", claims)
			return next(c)
		}
	}
}
