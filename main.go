package main

import (
	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/users"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/v1")

	users.SetRoutes(v1)

	app.Listen(":3001")
}
