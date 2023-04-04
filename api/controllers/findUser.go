package controllers

import (
	"backend/api/models"
	"backend/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindUser(c *fiber.Ctx) error {
	var users []models.User

	type UserFind struct {
		NameUser string `json:"nameUser"`
	}
	var userFind UserFind
	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")

	if err := c.BodyParser(&userFind); err != nil {
		return c.JSON(fiber.Map{
			"res": "error entrando data",
			"Err": err,
		})
	}
	fmt.Println(userFind, "-------------")
	findUsers := bson.D{{
		Key:   "nameUser",
		Value: userFind,
	}}
	cursor, err := db.Find(context.TODO(), findUsers, options.Find().SetLimit(4))
	if err != nil {
		return c.JSON(fiber.Map{
			"res": "error",
		})
	}
	for cursor.Next(context.TODO()) {
		fmt.Println(cursor, "+++++++")
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return c.JSON(fiber.Map{
		"res": "aaa",
	})
}
