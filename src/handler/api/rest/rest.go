package rest

import (
	"context"

	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/labstack/echo/v4"
)

type Rest struct {
	cfg    *config.AppConfig
	echo   *echo.Echo
	client *generated.Client
}

func Run(client *generated.Client, cfg *config.AppConfig, echo *echo.Echo, ctx context.Context) {
	rest := Rest{
		cfg:    cfg,
		echo:   echo,
		client: client,
	}

	NewAuth(&rest, ctx)
	NewRole(&rest, ctx)
	NewUser(&rest, ctx)
}
