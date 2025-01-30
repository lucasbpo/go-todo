package users

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	users := r.Group("/users")

	users.Post("/", addUser)
	users.Get("/", getAll)
	users.Get("/:id", getByID)
	users.Put("/:id", updateUser)
}
