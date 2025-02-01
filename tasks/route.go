package tasks

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	tasks := r.Group("/tasks")

	tasks.Post("/", addTask)
	tasks.Get("/", getAll)
	tasks.Get("/:id", getByID)
	tasks.Put("/:id", updateTask)
	tasks.Delete("/:id", deleteTask)
}
