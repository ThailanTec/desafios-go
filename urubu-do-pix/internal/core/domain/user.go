package domain

type User struct {
	ID       int     `json:"id" `
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required,email"`
	CPF      string  `json:"cpf" binding:"required"`
	Birthday string  `json:"birthday" binding:"required"`
	MoneyCod int     `json:"moneyCod"`
	SMoney   float64 `json:"sMoney"`
}
