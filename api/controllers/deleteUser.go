package controllers

import (
	"backend/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteUser(c *fiber.Ctx) error {
	client := database.DB()
	dataMiddleware := c.Context().UserValue("nameUser")
	fmt.Println(dataMiddleware)
	db := client.Database("goMoongodb").Collection("users")

	filter := bson.D{{Key: "nameUser", Value: dataMiddleware}}

	result, err := db.DeleteOne(context.TODO(), filter)
	if err != nil {
		return c.JSON(&fiber.Map{
			"res": "err server",
		})
	}
	if result.DeletedCount == 0 {
		return c.JSON(&fiber.Map{
			"res": "no se elimino ningun elemento",
		})
	} else {
		return c.JSON(&fiber.Map{
			"res": "user delete",
		})
	}
}
