package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ThailanTec/desafio-market/db"
	"github.com/ThailanTec/desafio-market/model"
)

func ValidaTransfer(uID1, valorTransfer int) (bool, error) {

	if GetSaldo(uID1) < valorTransfer || valorTransfer <= 0 {
		return false, errors.New("houve um erro na transacao valor insuficiente ou menor que 0")
	}

	return true, nil

}

func DebValor(uID1, valorTransfer int) (bool, int) {

	var user1 = model.User{ID: uID1}
	db.DB.First(&user1)

	saldo1 := GetSaldo(uID1)

	result1 := saldo1 - valorTransfer
	user1.Saldo = result1
	db.DB.Save(user1)

	fmt.Println("VALOR FINAL", result1, valorTransfer)

	return true, result1
}

func CredValor(uID2, valorReceb int) (bool, int) {

	var user1 = model.User{ID: uID2}
	db.DB.First(&user1)

	saldo1 := GetSaldo(uID2)

	result2 := saldo1 + valorReceb
	user1.Saldo = result2
	db.DB.Save(user1)

	fmt.Println("VALOR FINAL", result2, valorReceb)

	return true, result2
}

func RollBack(uID1 int) {

	var user1 = model.User{ID: uID1}
	db.DB.First(&user1)

	valor := GetSaldo(uID1)

	str := strconv.Itoa(valor)
	db.DB.SavePoint(str)

	db.DB.RollbackTo(str)
}

func GetSaldo(id int) int {
	var user = model.User{ID: id}
	db.DB.First(&user)

	return user.Saldo
}

func DataTransfer(debID, CredID, valorDeb int) int {

	user1 := model.User{ID: debID}
	db.DB.First(&user1)

	user2 := model.User{ID: CredID}
	db.DB.First(&user2)

	valorUser1 := valorDeb

	return valorUser1

}
