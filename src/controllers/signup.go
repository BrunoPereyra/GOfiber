package controllers

import (
	"backend/src/database"
	"backend/src/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var user models.User
	var userExist models.User

	if err := c.BodyParser(&user); err != nil {
		return c.JSON(&fiber.Map{
			"err": "valito res",
		})
	}

	if err := user.Validate(); err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	client := database.DB()
	db := client.Database("goMoongodb").Collection("users")
	filter := bson.D{{Key: "nameUser", Value: user.NameUser}}

	if err := db.FindOne(context.TODO(), filter).Decode(&userExist); err != nil {

		if err == mongo.ErrNoDocuments {

			passwordHash, err := HashPassword(user.Password)
			if err != nil {
				return c.JSON(&fiber.Map{
					"error":     "error in HashPassword",
					"dataError": err,
				})
			}
			newUser := []interface{}{
				bson.M{
					"nameUser":  user.NameUser,
					"password":  passwordHash,
					"createdAt": time.Now(),
					"updatedAt": time.Now(),
				},
			}

			result, err := db.InsertMany(context.TODO(), newUser)
			if err != nil {
				return c.JSON(&fiber.Map{
					"error": "error al cargar los usuarios",
				})
			} else {
				return c.JSON(&fiber.Map{
					"res": result,
				})
			}

		} else {
			return c.JSON(&fiber.Map{
				"error": "error",
			})
		}
	}

	return c.JSON(&fiber.Map{
		"data": "userExist",
	})

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
