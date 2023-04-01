package controllers

import (
	"backend/src/database"
	"backend/src/models"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	c.BodyParser(&user)
	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")

	result, err := db.InsertOne(context.TODO(), bson.D{{
		Key:   "name",
		Value: user.NameUser,
	}})
	fmt.Print(user)
	if err != nil {
		return c.JSON(&fiber.Map{
			"data": err,
		})
	}
	return c.JSON(&fiber.Map{
		"data": result,
	})
}
