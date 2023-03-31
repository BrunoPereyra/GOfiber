package main

import (
	"backend/backend/models"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Definimos la estructura de la tar
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	// Configuración de la aplicación Fiber
	app := fiber.New()

	// Configuración de la conexión con la base de datos MongoDB
	clientOptions := options.Client().ApplyURI("mongodb+srv://bruno:Ej622XsRFaiuh9kg@cluster0.0leyfts.mongodb.net/TDD?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	db := client.Database("goMoongodb").Collection("users")

	// Ruta para crear una tarea
	app.Post("/user", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)
		result, err := db.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
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
	})

	// Iniciamos la aplicación en el puerto 3000
	app.Listen(":3000")
}
