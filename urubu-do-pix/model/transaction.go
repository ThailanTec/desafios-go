package model

type DUser struct {
	TransactionAmount int    `json:"transaction_amount"`
	Description       string `json:"description"`
	PaymentMethodID   string `json:"payment_method_id"`
	Payer             Payer  `json:"payer"`
}
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}
type Payer struct {
	Email          string         `json:"email"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Identification Identification `json:"identification"`
}

type TransData struct {
	TransactionDat TransactionData `json:"transaction_data"`
}
type TransactionData struct {
	Qr_Code string `json:"qr_code"`
}
