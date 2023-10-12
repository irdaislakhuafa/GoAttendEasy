package scheduller

import (
	"context"

	"github.com/go-co-op/gocron"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/smtp"
)

type Scheduller struct {
	client *generated.Client
	smtp   smtp.GoMailInterface
	cron   *gocron.Scheduler
	cfg    *config.AppConfig
}

func InitScheduller(client *generated.Client, smtp smtp.GoMailInterface, cron *gocron.Scheduler, cfg *config.AppConfig, ctx context.Context) {
	scheduller := &Scheduller{
		client: client,
		smtp:   smtp,
		cron:   cron,
		cfg:    cfg,
	}

	NewReminder(scheduller, ctx)
	scheduller.cron.StartAsync()
}
