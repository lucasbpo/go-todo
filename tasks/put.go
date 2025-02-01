package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func updateTask(c *fiber.Ctx) error {

	body := new(Task)
	result := Task{}

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	err := db.UpdateByID("tasks", c.Params("id"), body, &result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(result)
}
