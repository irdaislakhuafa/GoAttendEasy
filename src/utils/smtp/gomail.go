package smtp

import (
	mail "github.com/go-mail/gomail"
	"github.com/irdaislakhuafa/GoAttendEasy/src/utils/config"
)

type GoMailInterface interface {
	DialAndSend(messages ...*mail.Message) error
}

type gomail struct {
	dialer *mail.Dialer
}

func InitGoMail(cfg config.SMTP) GoMailInterface {
	result := gomail{
		dialer: mail.NewDialer(cfg.Host, int(cfg.Port), cfg.Username, cfg.Password),
	}

	return &result
}

func (g *gomail) DialAndSend(messages ...*mail.Message) error {
	if err := g.dialer.DialAndSend(messages...); err != nil {
	}
	return nil
}
