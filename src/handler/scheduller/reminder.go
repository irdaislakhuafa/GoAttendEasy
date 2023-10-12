package scheduller

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	mail "github.com/go-mail/gomail"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/reminder"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/role"
	"github.com/irdaislakhuafa/GoAttendEasy/src/schema/generated/user"
)

type ReminderInterface interface {
	ClockInOutReminder(ctx context.Context)
}

type schedullerReminder struct {
	scheduller *Scheduller
	day        map[int8]string
	reminder   *generated.Reminder
}

func NewReminder(scheduller *Scheduller, ctx context.Context) ReminderInterface {
	reminder := &schedullerReminder{
		scheduller: scheduller,
		day: map[int8]string{
			1: "Sunday",
			2: "Monday",
			3: "Tuesday",
			4: "Wednesday",
			5: "Thursday",
			6: "Friday",
			7: "Saturday",
		},
	}

	reminder.ClockInOutReminder(ctx)
	return reminder
}

func (r *schedullerReminder) ClockInOutReminder(ctx context.Context) {
	r.scheduller.cron.Every(1).Day().At("00:00:00").Do(func() {
		log.Println("set notification for clock in and out from reminder configuration")
		now := time.Now()
		weekDay := (now.Weekday() + 1)

		reminder, err := r.scheduller.client.Reminder.Query().Where(reminder.Day(int(weekDay)), reminder.IsDeleted(false)).First(ctx)
		if err != nil {
			log.Printf("clockin reminder at %v not found, %v\n", r.day[int8(weekDay)], err)
			r.reminder = nil
			return
		}

		role, err := r.scheduller.client.Role.Query().Where(role.NameEqualFold("employee"), role.IsDeleted(false)).First(ctx)
		if err != nil {
			log.Printf("error while getting role, %v\n", err)
			return
		}

		listUser, err := r.scheduller.client.User.Query().Where(user.RoleID(role.ID), user.IsDeleted(false)).All(ctx)
		if err != nil {
			log.Printf("failed to get list user with role employee, %v\n", err)
			return
		}

		scheduller := gocron.NewScheduler(time.Local)

		// clock in reminder
		if !reminder.In.IsZero() {
			scheduller.Every(1).Day().At(reminder.In.Format("15:04:05")).Do(func() {
				listMsg := []*mail.Message{}
				for _, u := range listUser {
					msg := mail.NewMessage(func(m *mail.Message) {
						m.SetHeaders(map[string][]string{
							"From":     r.scheduller.cfg.SMTP.Information.Senders,
							"Reply-To": r.scheduller.cfg.SMTP.Information.Senders,
							"To":       {u.Email},
							"Subject":  {"Clock In Reminder"},
						})
						m.SetBody("text/html", fmt.Sprintf("<center><h1>Hi %v, Clock In Today is %v</h1></center>", u.Name, reminder.In.Format("15:04")))
					})
					listMsg = append(listMsg, msg)
				}

				if err := r.scheduller.smtp.DialAndSend(listMsg...); err != nil {
					log.Printf("failed to send emails, %v\n", err)
					return
				}
			})
		}

		// clock out reminder
		if !reminder.Out.IsZero() {
			scheduller.Every(1).Day().At(reminder.Out.Format("15:04:05")).Do(func() {
				listMsg := []*mail.Message{}
				for _, u := range listUser {
					msg := mail.NewMessage(func(m *mail.Message) {
						m.SetHeaders(map[string][]string{
							"From":     r.scheduller.cfg.SMTP.Information.Senders,
							"Reply-To": r.scheduller.cfg.SMTP.Information.Senders,
							"To":       {u.Email},
							"Subject":  {"Clock out Reminder"},
						})
						m.SetBody("text/html", fmt.Sprintf("<center><h1>Hi %v, Clock Out Today is %v</h1></center>", u.Name, reminder.In.Format("15:04")))
					})
					listMsg = append(listMsg, msg)
				}

				if err := r.scheduller.smtp.DialAndSend(listMsg...); err != nil {
					log.Printf("failed to send emails, %v\n", err)
					return
				}
			})
		}

		scheduller.StartBlocking()
		log.Println("done send reminder notification")
	})
}
