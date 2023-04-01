package controllers

import (
	"backend/backend/database"
	"backend/backend/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Signup(c *fiber.Ctx) error {

	var user models.User
	var result models.User

	if err := c.BodyParser(user); err != nil {
		return err
	}

	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")

	filter := bson.D{{"name", user.Name}}
	err := db.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		return c.JSON(&fiber.Map{
			"error": err,
		})
	}
	return c.JSON(&fiber.Map{
		"data": result,
	})

}
