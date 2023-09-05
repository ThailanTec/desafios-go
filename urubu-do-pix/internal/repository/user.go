package repository

import (
	"errors"

	"github.com/ThailanTec/urubu-do-pix/internal/core/domain"
)

type User struct {
	User domain.User
}

func NewUser(user domain.User) *User {
	return &User{User: user}
}

func (u *User) GetUser(email string) (*domain.User, error) {

	if u.User.Email == "" {
		return nil, errors.New("email not exist")
	}

	return &domain.User{}, nil
}
