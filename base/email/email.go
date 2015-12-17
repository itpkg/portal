package email

type Provider interface {
	Send(to []string, subject, body string, attach ...string)
}
