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

func GetById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var user1 = model.User{ID: id}
	db.DB.First(&user1)

	c.JSON(http.StatusOK, user1)
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

	valorTransfer := DataTransfer(uID1, uID2, 10)

	b, err := ValidaTransfer(uID1, valorTransfer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Erro ao realizar transferencia")
		return
	} else if !b {
		RollBack(uID1)
	} else if b {
		DebValor(uID1, valorTransfer)

		CredValor(uID2, valorTransfer)
	}

	c.JSON(http.StatusOK, "Transferencia realizada com sucesso.")
}
