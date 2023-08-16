package server

import (
	"github.com/ThailanTec/urubu-do-pix/controller"
	"github.com/ThailanTec/urubu-do-pix/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	db.Connect()
	router := gin.Default()

	// Routes
	router.GET("/", controller.BVindo)
	router.POST("/createUser", controller.CreateUser)
	router.GET("/getAll", controller.GetAllUser)
	return router
}
