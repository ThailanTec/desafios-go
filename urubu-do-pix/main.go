package main

import "github.com/ThailanTec/urubu-do-pix/server"

func main() {

	r := server.SetupRouter()
	r.Run(":8001")

}
