package routes

import (
	"backend/api/controllers"
	"backend/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	app.Post("/login", controllers.Login)
	app.Get("/user", middleware.MiddlewareUseExtractor(), controllers.GeAllUser)
	app.Post("/signup", controllers.Signup)

	app.Get("/delete", middleware.MiddlewareUseExtractor(), controllers.DeleteUser)

}
