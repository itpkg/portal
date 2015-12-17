package email

import (
	"github.com/op/go-logging"
)

type LocalProvider struct {
	Logger *logging.Logger `inject:""`
}

func (p *LocalProvider) Send(to []string, subject, body string, attach ...string) {
	p.Logger.Debug("SENDEMAIL\nTO: %v\nSUBJECT:\n%s\nBODY:\n%s", to, subject, body)
}
