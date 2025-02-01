package tasks

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func deleteTask(c *fiber.Ctx) error {
	err := db.DeleteByID("tasks", c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusNoContent).SendString("")
}
