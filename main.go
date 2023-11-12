package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"tz/internal/routes"
	"tz/pkg/database"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	db := database.InitDB()

	if db == nil {
		log.Fatal("Failed to initialize the database")
	}

	routes.Requests(app, db)
	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
