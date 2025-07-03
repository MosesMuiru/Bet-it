package main

import (
	db "kwik/storage"
	"kwik/storage/models"
)

func main() {

	db.ConnectDb()

	DB := db.DB()

	// autho migrate models

	DB.AutoMigrate(&models.Slip{})
	DB.AutoMigrate(&models.Selection{})

}
