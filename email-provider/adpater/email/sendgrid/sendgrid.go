package adapter

import (
	"github.com/ThailanTec/desafios-go/email-provider/internal/core/domain"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendGridMail struct {
	c      *sendgrid.Client
	sender *mail.Email
}

func NewSendGrid(settings domain.EmailSettings) *sendGridMail {
	return &sendGridMail{
		c: sendgrid.NewSendClient(settings.APIKey), sender: mail.NewEmail(settings.SenderName, settings.SenderEmail),
	}
}

func (s *sendGridMail) Send(e domain.Email) (err error) {
	message := mail.NewSingleEmail(s.sender, e.Subject, mail.NewEmail(e.Receiver, e.To), "", e.Content)
	_, err = s.c.Send(message)
	return
}
