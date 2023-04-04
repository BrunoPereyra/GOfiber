package routes

import (
	"backend/api/controllers"
	"backend/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	app.Post("/signup", controllers.Signup)
	app.Post("/login", controllers.Login)
	app.Post("/findUser", middleware.MiddlewareUseExtractor(), controllers.FindUser)

	app.Get("/user", middleware.MiddlewareUseExtractor(), controllers.GeAllUser)
	app.Get("/delete", middleware.MiddlewareUseExtractor(), controllers.DeleteUser)

}
