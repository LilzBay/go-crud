package main

import (
	"log"

	"github.com/LilzBay/go-crud/initializers"
	"github.com/LilzBay/go-crud/models"
	"github.com/joho/godotenv"
)

func init() {
	// initializers.LoadEnvVariables()
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
