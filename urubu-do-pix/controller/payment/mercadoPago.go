package payment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ThailanTec/urubu-do-pix/model"
)

func CreatePayment(usr model.User) (b []byte, e error) {

	url := "https://api.mercadopago.com/v1/payments"

	data := TransData(usr)

	tData, err := json.Marshal(data)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(tData))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer TEST-5232187670379919-081619-5f1fa4af5fd8cb684f1787106f78817f-239690431")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Handle response
	return body, nil
}

func TransData(user model.User) (dataT model.DUser) {

	data := model.DUser{
		TransactionAmount: int(user.SMoney),
		Description:       "Urubu do PIX",
		PaymentMethodID:   "pix",
		Payer: model.Payer{
			Email:     user.Email,
			FirstName: user.Name,
			LastName:  user.Name,
			Identification: model.Identification{
				Type:   "CPF",
				Number: user.CPF,
			},
		},
	}

	return data

}
