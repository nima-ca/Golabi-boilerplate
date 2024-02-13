package db

import (
	"Golabi-boilerplate/models"
	"fmt"
	"log"
)

func MigrateDB() {

	// Connect to DB
	Connect()

	// Auto migrate models
	err := DB.AutoMigrate(
		&models.User{},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration was successful")
}
