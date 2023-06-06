package service

import (
	"github.com/ThailanTec/desafios-go/email-provider/internal/core/domain"
	"github.com/ThailanTec/desafios-go/email-provider/internal/core/ports"
)

type EmailServiceImplementation struct {
	emailRepository ports.EmailRepository
}

// NewEmailService cria uma nova instância do serviço de e-mail
func NewEmailService(emailRepository ports.EmailRepository) *EmailServiceImplementation {
	return &EmailServiceImplementation{
		emailRepository: emailRepository,
	}
}

// SendEmail envia um e-mail usando o Mailjet
func (s *EmailServiceImplementation) Send(e domain.Email) (err error) {
	// Lógica para configurar o cliente Mailjet e enviar o e-mail
	// ...

	// Exemplo de uso do Mailjet para enviar o e-mail
	mailjetClient := mailjet.NewMailjetClient(apiKey, apiSecret)
	sendEmailRequest := &mailjet.SendMailRequest{
		FromEmail: "sender@example.com",
		FromName:  "Sender",
		Recipients: []mailjet.Recipient{
			{
				Email: "recipient@example.com",
			},
		},
		Subject:  e.Subject,
		TextPart: e.Content,
	}
	_, err := mailjetClient.SendMail(sendEmailRequest)
	if err != nil {
		return err
	}

	// Salva o e-mail enviado no repositório
	err = s.emailRepository.SaveEmail(email)
	if err != nil {
		// Lidar com erros ao salvar o e-mail no repositório
		return err
	}

	return nil
}
