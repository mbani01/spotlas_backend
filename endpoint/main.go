package main

import (
	"endpoint/database"
	"endpoint/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
	database.InitConnection()
	app.Get("/", handlers.ValidateSpotParam, handlers.GetSpots)
    app.Listen(":3000")
}