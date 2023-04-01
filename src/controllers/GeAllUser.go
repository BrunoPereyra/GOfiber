package controllers

import (
	"backend/src/database"
	"backend/src/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GeAllUser(c *fiber.Ctx) error {
	var users []models.User

	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")

	result, err := db.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.JSON(&fiber.Map{
			"error": err,
		})
	}
	for result.Next(context.TODO()) {
		var user models.User
		result.Decode(&user)
		users = append(users, user)
	}
	return c.JSON(&fiber.Map{
		"user": users,
	})
}
