package controllers

import (
	"time"

	"flutter-dev.info/go-backend/database"
	_ "flutter-dev.info/go-backend/database"
	"flutter-dev.info/go-backend/models"
	"flutter-dev.info/go-backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	data := fiber.Map{}

	if err := c.BodyParser(&data); err != nil {
		c.Status(500)
		return err
	}

	if data["password"] != data["cpassword"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password do not matched",
		})
	}
	if data["email"] == "" || data["first_name"] == "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Provide All details",
		})
	}
	var count int64
	database.DB.Model(&models.User{}).Where("email", data["email"].(string)).Count(&count)

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"].(string)), 14)

	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Password:  password,
		Email:     data["email"].(string),
		UserType:  data["user_type"].(string),
		Phone:     data["phone"].(string),
	}

	if count != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email id already used",
		})
	}
	result := database.DB.Create(&user)
	if result.RowsAffected == 0 {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": result.Error,
		})
	}
	return c.JSON(fiber.Map{"data": user})
}

func Login(c *fiber.Ctx) error {
	data := fiber.Map{}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Model(&models.User{}).Where("email = ? ", data["email"].(string)).First(&user)
	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "User not found with this email ID",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"].(string))); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	token, err := utils.GenerateJwt(user.Id)
	if err != nil {
		c.Status(500)
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	id, _ := utils.ParseJwt(token)
	var user_data models.User
	database.DB.Model(user).Where("id = ?", id).First(&user_data)
	return c.JSON(fiber.Map{
		"message":   "Login Success",
		"user_data": user_data,
		"token":     token,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	id, _ := utils.ParseJwt(cookie)
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(fiber.Map{"user": user})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "Logout Success",
	})
}
