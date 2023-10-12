package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewUserWeb(web *Web) {
	web.echo.GET("/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
}
