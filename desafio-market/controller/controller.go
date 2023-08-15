package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ThailanTec/desafio-market/db"
	"github.com/ThailanTec/desafio-market/model"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var user []model.User
	if result := db.DB.Find(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	c.JSON(http.StatusOK, &user)
}

func Transfer(c *gin.Context) {

	uID1, _ := strconv.Atoi(c.Param("id"))
	uID2, _ := strconv.Atoi(c.Param("id2"))

	// Dados transferencia user1
	var user1 = model.User{ID: uID1}
	db.DB.First(&user1)
	// Dados transferencia user2
	var user2 = model.User{ID: uID2}
	db.DB.First(&user2)

	valorTransfer := DataTransfer(uID1, uID2, 60)

	if GetSaldo(uID1) < valorTransfer || valorTransfer <= 0 {
		c.JSON(500, "Saldo Insuficiente")
		return
	} else {
		fmt.Println(user1)
		fmt.Println(user2)

		saldo1 := GetSaldo(uID1)
		saldo2 := GetSaldo(uID2)

		result1 := saldo1 - valorTransfer
		user1.Saldo = result1
		db.DB.Save(user1)

		result2 := saldo2 + valorTransfer
		user2.Saldo = result2
		db.DB.Save(user2)

		fmt.Println(result1)
		fmt.Println(result2)
	}

	c.JSON(http.StatusOK, "Transferencia realizada com sucesso.")
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
