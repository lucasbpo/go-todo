package main

import (
	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/tags"
	"lucasbpo.com.br/lucasbpo/todo/tasks"
	"lucasbpo.com.br/lucasbpo/todo/users"
)

func main() {
	app := fiber.New()

	v1 := app.Group("/api/v1")

	users.SetRoutes(v1)
	tags.SetRoutes(v1)
	tasks.SetRoutes(v1)

	app.Listen(":3001")
}
