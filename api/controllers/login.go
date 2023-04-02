package controllers

import (
	"backend/api/models"
	"backend/database"
	"backend/jwt"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
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

	err := db.FindOne(context.TODO(), filter).Decode(&userExist)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(&fiber.Map{
				"res": "user no existe",
			})
		}
		return c.JSON(&fiber.Map{
			"res": "error server",
			"err": err,
		})
	}
	if err := comparePassword([]byte(userExist.Password), []byte(user.Password)); err != nil {
		return c.JSON(&fiber.Map{
			"res": "password incorrect",
		})
	}

	signedToken, err := jwt.TokenCreate(user)
	return c.JSON(&fiber.Map{
		"user":  user,
		"Token": signedToken,
	})
}

func comparePassword(hash []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err
}
