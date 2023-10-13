package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewDocs(web *Web) {
	web.echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf(`
		<center>
			<h1>Read The Documentation in Postman <a href="%s">Here</a></h1>
		</center>
		`, web.cfg.App.Documentation.Postman))
	})
}
