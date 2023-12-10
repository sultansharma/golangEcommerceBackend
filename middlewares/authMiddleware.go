package middlewares

import (
	"flutter-dev.info/go-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	//cookie := c.Cookies("jwt") // jwt token from cookie if request from web
	var headers = c.GetReqHeaders() // jwt token from headerif req from API
	_, err := utils.ParseJwt(headers["Token"][0])

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{"message": "Unauthorized"})
	}
	return c.Next()
}
