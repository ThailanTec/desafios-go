package model

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	CPF      int64  `json:"cpf"`
	Birthday string `json:"birthday"`
	Money    int    `json:"money"`
}
