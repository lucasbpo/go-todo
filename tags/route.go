package tags

import "github.com/gofiber/fiber/v2"

func SetRoutes(r fiber.Router) {
	tags := r.Group("/tags")

	tags.Post("/", addTag)
	tags.Get("/", getAll)
	tags.Get("/:id", getByID)
	tags.Put("/:id", updateTag)
	tags.Delete("/:id", deleteTag)
}
