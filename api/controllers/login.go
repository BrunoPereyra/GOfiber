package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// client := database.DB()
	// db := client.Database("goMoongodb").Collection("users")
	authHeader := c.Get("Authorization")
	token := strings.Replace(authHeader, "Bearer ", "", 1)
	// jwt.TokenVerify(token)
	fmt.Println(token)
	return c.JSON(&fiber.Map{
		"user": "A",
	})
}
