package main

import (
	"backend/api/routes"
	"backend/jwt"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Definimos la estructura de la tar
type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	// Configuración de la aplicación Fiber
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODA0OTQ0OTYsImlkIjoiMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwidXNlcm5hbWUiOiJuVWU5ciJ9.m6TZgNhWRcQO5blDJTkiaiNYEBC21F9EchJQ9feslLM"
	fmt.Println("_____")
	nameUser, _ := jwt.ExtractDataFromToken(tokenString)
	fmt.Println(nameUser + "AA")
	fmt.Println("_____")

	app := fiber.New()
	routes.UserRoute(app)

	app.Listen(":3000")
}
