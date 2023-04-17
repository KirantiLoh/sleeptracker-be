package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kirantiloh/dashtrack-be/pkg/routes"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")

	routes.InitializeRoutes(v1)

	app.Listen(":8080")
}
