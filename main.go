package main

import (
	"github.com/gofiber/fiber/v2"
	"tz/internal/routes"
)

func main() {
	app := fiber.New()
	routes.Requests(app)

	err := app.Listen(`:8080`)
	if err != nil {
		return
	}
}
