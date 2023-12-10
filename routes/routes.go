package routes

import (
	"flutter-dev.info/go-backend/controllers"
	"flutter-dev.info/go-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.HomePage)
	app.Get("/users", controllers.Users)
	app.Post("/auth/register", controllers.Register)
	app.Post("/auth/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated) //Midlleware

	app.Post("/auth/logout", controllers.Logout)
	app.Get("/user/info", controllers.User)
}
