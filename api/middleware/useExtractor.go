package middleware

import (
	"backend/jwt"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareUseExtractor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		nameUser, err := jwt.ExtractDataFromToken(token)
		if err != nil {
			fmt.Println(err)
			return c.JSON(&fiber.Map{
				"res": "token error or invalid",
				"err": err,
			})
		}
		c.Context().SetUserValue("nameUser", nameUser)
		return c.Next()
	}
}
