package server

import (
	"github.com/ThailanTec/desafio-market/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Routes
	router.GET("/getAll", controller.GetAllUser)
	router.GET("/getById/:id", controller.GetById)
	router.GET("/transfer/:id/:id2", controller.Transfer)

	return router
}
