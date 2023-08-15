package main

import (
	"github.com/ThailanTec/desafio-market/db"
	"github.com/ThailanTec/desafio-market/server"
)

func main() {
	db.Connect()
	r := server.SetupRouter()
	r.Run(":8001")
}
