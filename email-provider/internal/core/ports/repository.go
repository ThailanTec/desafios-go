package ports

import "github.com/ThailanTec/desafios-go/email-provider/internal/core/domain"

type Email interface {
	Send(e domain.Email) (err error)
}
