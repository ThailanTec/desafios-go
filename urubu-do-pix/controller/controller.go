package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/ThailanTec/urubu-do-pix/controller/payment"
	"github.com/ThailanTec/urubu-do-pix/db"
	"github.com/ThailanTec/urubu-do-pix/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
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

	if GetEmail(nUser.Email) {
		c.JSON(400, "E-mail já existe. Manda outro")
		return
	}

	usr, err := payment.CreatePayment(nUser)
	if err != nil {
		return
	}
	var dF model.TransData

	if err := json.Unmarshal(usr, &dF); err != nil {
		c.JSON(400, err)
		return
	}

	// Paramos aqui, receber os dados do pix que vem do MP.

	cod := rand.Intn(1000)

	nUser.ID = cod
	nUser.MoneyCod = cod

	db.DB.Create(&nUser)
	c.JSON(http.StatusCreated, nUser)
}

func GetAllUser(c *gin.Context) {

	users := []model.User{}

	db.DB.Find(&users)

	c.JSON(200, &users)
}

func GetEmail(email string) bool {
	var user model.User

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false // E-mail não encontrado no banco de dados
		}
	}

	return true
}
