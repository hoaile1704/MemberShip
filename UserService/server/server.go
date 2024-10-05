package server

import (
	"UserService/internal/middleware"
	"UserService/internal/routers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func Run() {
	userHandler := InitializeUserHandler()
	app := fiber.New()
	app.Use(healthcheck.New())
	app.Use(middleware.ResponseMiddleware)
	routers.UserRouter(app, userHandler)
	log.Fatal(app.Listen(":3000"))
}
