package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-playground/validator/v10"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/api/rest"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/scheduller"
	"github.com/irdaislakhuafa/GoAttendEasy/src/handler/web"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/connection"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/customvalidator"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/flags"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/smtp"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	cfgPath     = "etc/cfg/"
	cfgFileName = "conf.json"
)

func main() {
	// get env from option, default is local
	env, err := flags.ParseFlags(cfgPath, cfgFileName)
	if err != nil {
		panic(err)
	}

	// read config from config file
	cfg, err := config.ReadConfigFromFile(fmt.Sprintf("%s/%s/%s", cfgPath, *env, cfgFileName))
	if err != nil {
		panic(err)
	}

	// open db connection
	client, err := connection.OpenConnection(cfg)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// init context
	ctx := context.Background()

	// run database migration
	if err := client.Schema.Create(ctx); err != nil {
		panic(err)
	}

	// initialize smtp
	smtp := smtp.InitGoMail(cfg.SMTP)

	// initialize go cron
	tz, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	cron := gocron.NewScheduler(tz)

	// initilize echo
	e := echo.New()

	// added customvalidator
	e.Validator = &customvalidator.CustomValidator{Validator: validator.New()}
	e.Use(middleware.Logger())

	// run web app
	web.Run(client, &cfg, e)

	// run rest api
	rest.Run(client, &cfg, e, ctx)

	// start scheduller
	scheduller.InitScheduller(client, smtp, cron, &cfg, ctx)

	// start echo http server
	e.Start(fmt.Sprintf(":%d", cfg.Server.Port))

}
