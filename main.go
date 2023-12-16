package main

import (
	"go-rest-api-crud/database"
	"go-rest-api-crud/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot Load Env")
	}

	database.ConnectDB()

	app := fiber.New()
	routes.SetUpRoutes(app)

	app.Listen(":8080")
}
