package model

type User struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Saldo int    `json:"saldo"`
}

type Transfer struct {
	US1   int `json:"id"`
	US2   int `json:"id2"`
	Saldo int `json:"money"`
}
