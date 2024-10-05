package routers

import (
	"UserService/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App, userHandler *handlers.UserHandler) {
	app.Get("/users/:id", userHandler.GetUserHandler)
	app.Post("/users", userHandler.CreateUserHandler)
}
