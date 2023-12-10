package main

import (
	"log"

	"flutter-dev.info/go-backend/database"
	"flutter-dev.info/go-backend/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	database.Connect()
	log.Fatal(app.Listen(":3000"))

}
