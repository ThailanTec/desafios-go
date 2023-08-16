package controller

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/ThailanTec/urubu-do-pix/db"
	"github.com/ThailanTec/urubu-do-pix/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BVindo(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo grande amigo! Vamos mudar de vida! Pode vir ao balcão!")
}

func CreateUser(c *gin.Context) {

	nUser := model.User{}

	if errA := c.ShouldBindBodyWith(&nUser, binding.JSON); errA != nil {
		c.JSON(http.StatusBadRequest, errA)
		return
	}

	nUser.ID = rand.Intn(1000)
	GetEmail(nUser.Email)

	db.DB.Create(&nUser)
	c.JSON(http.StatusOK, nUser)
}

func GetAllUser(c *gin.Context) {

	users := []model.User{}

	db.DB.Find(&users)

	c.JSON(200, &users)
}

func GetEmail(email string) {

	userE := model.User{Email: email}
	db.DB.Find(&userE)

	fmt.Println(userE)
}
