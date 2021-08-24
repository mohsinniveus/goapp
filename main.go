package main

import (
	"goapp/models"
	"goapp/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})

	r := routes.SetupRoutes(db)
	r.Run(":7070")
}
