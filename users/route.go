package users

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	users := r.Group("/users")

	users.Post("/", addUser)
}
