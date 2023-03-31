package main

import (
	"backend/backend/routes"

	"github.com/gofiber/fiber/v2"
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
	// Ruta para crear una tarea
	routes.UserRoute(app)

	// app.Get("/user", func(c *fiber.Ctx) error {
	// 	var users []models.User
	// 	result, err := db.Find(context.TODO(), bson.M{})
	// 	if err != nil {
	// 		return c.JSON(&fiber.Map{
	// 			"error": err,
	// 		})
	// 	}
	// 	for result.Next(context.TODO()) {
	// 		var user models.User
	// 		result.Decode(&user)
	// 		users = append(users, user)
	// 	}
	// 	return c.JSON(&fiber.Map{
	// 		"user": users,
	// 	})
	// })
	// Iniciamos la aplicación en el puerto 3000
	app.Listen(":3000")
}
