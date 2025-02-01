package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func getAll(c *fiber.Ctx) error {
	documents := []Task{}

	err := db.Find("tasks", &documents)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error)
	}

	return c.JSON(documents)
}

func getByID(c *fiber.Ctx) error {
	document := Task{}

	err := db.FindByID("tasks", c.Params("id"), &document)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(document)
}
