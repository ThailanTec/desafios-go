package model

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Saldo int    `json:"saldo"`
}
