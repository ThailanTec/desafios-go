package ports

import "github.com/ThailanTec/urubu-do-pix/internal/core/domain"

type UserService interface {
	CreateUser(user domain.User)
	GetAllUser() domain.User
}
