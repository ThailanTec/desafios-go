package ports

import "github.com/ThailanTec/urubu-do-pix/internal/core/domain"

type UserRepo interface {
	GetUser(email string) (domain.User, error)
	GetEmail(email string) bool
	GetAllUser() domain.User
}
