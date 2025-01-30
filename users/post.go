package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func addUser(c *fiber.Ctx) error {
	body := new(User)

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	id, err := db.Insert("users", body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	body.ID = id

	return c.Status(http.StatusCreated).JSON(body)
}
