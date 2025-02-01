package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func addTask(c *fiber.Ctx) error {
	body := new(Task)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("tasks", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id

	return c.Status(http.StatusCreated).JSON(body)
}
