package main

import (
	"github.com/famdar/database"
	"github.com/famdar/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()
	r.Run(":8080")
}
