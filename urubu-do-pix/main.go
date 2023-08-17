package main

import (
	"github.com/ThailanTec/urubu-do-pix/db"
	"github.com/ThailanTec/urubu-do-pix/server"
)

func main() {

	db.Connect()

	r := server.SetupRouter()
	r.Run(":8001")

}
