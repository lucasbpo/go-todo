package tags

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func getAll(c *fiber.Ctx) error {
	documents := []Tag{}

	err := db.Find("tags", &documents)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error)
	}

	return c.JSON(documents)
}

func getByID(c *fiber.Ctx) error {
	document := Tag{}

	err := db.FindByID("tags", c.Params("id"), &document)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}
