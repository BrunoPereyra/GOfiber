package main

import (
	"backend/api/routes"

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

	app.Listen(":3000")
}
