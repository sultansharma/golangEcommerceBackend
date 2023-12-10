package controllers

import (
	_ "fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"Page": "Home", "path": os.Getenv("GOPATH")})
}

func Users(c *fiber.Ctx) error {
	var data fiber.Map

	if err := c.BodyParser(&data); err != nil {
		return c.JSON("No Data" + err.Error())
	}
	return c.JSON(data["songs"])
}
