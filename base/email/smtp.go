package email

import (
	"github.com/op/go-logging"
	"gopkg.in/gomail.v2"
)

type Smtp struct {
	Host     string
	Port     int
	From     string
	Username string
	Password string
	Ssl      bool
}

type SmtpProvider struct {
	Func   func() (*Smtp, error)
	Logger *logging.Logger `inject:""`
}

func (p *SmtpProvider) Send(to []string, subject, body string, attach ...string) {

	go func() {
		smtp, err := p.Func()
		if err != nil {
			p.Logger.Error("bad in sendmail: %v", err)
			return
		}

		msg := gomail.NewMessage()
		msg.SetHeader("From", smtp.From)
		msg.SetHeader("To", to...)
		msg.SetHeader("Subject", subject)
		msg.SetBody("text/html", body)
		for _, a := range attach {
			msg.Attach(a)
		}

		dialer := gomail.NewPlainDialer(smtp.Host, smtp.Port, smtp.Username, smtp.Password)
		if err := dialer.DialAndSend(msg); err != nil {
			p.Logger.Error("bad in sendmail: %v", err)
		}
	}()

}
