package web

import (
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/labstack/echo/v4"
)

type Web struct {
	client *generated.Client
	cfg    *config.AppConfig
	echo   *echo.Echo
}

func Run(client *generated.Client, cfg *config.AppConfig, echo *echo.Echo) {
	result := Web{
		client: client,
		cfg:    cfg,
		echo:   echo,
	}

	NewUserWeb(&result)
}

func (r *Web) GetEcho() *echo.Echo {
	return r.echo
}
