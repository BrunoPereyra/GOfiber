package controllers

import (
	"backend/api/models"
	"backend/api/validator"
	"backend/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindUser(c *fiber.Ctx) error {
	var users []models.User

	var userFind validator.UserFind
	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")

	if err := c.BodyParser(&userFind); err != nil {
		return c.JSON(fiber.Map{
			"res": "error entrando data",
			"Err": err,
		})
	}
	if err := userFind.ValidateUserFind(); err != nil {
		return c.JSON(fiber.Map{
			"res": err,
		})
	}
	findUsers := bson.D{{
		Key:   "nameUser",
		Value: userFind.NameUser,
	}}
	cursor, err := db.Find(context.TODO(), findUsers, options.Find().SetLimit(4))
	if err != nil {
		return c.JSON(fiber.Map{
			"res": "error",
		})
	}
	for cursor.Next(context.TODO()) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	fmt.Println(users[0].NameUser, "YYYYYY")
	return c.JSON(fiber.Map{
		"res": users,
	})
}
