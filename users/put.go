package users

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"lucasbpo.com.br/lucasbpo/todo/db"
)

func updateUser(c *fiber.Ctx) error {

	body := new(User)
	result := User{}

	if err := c.BodyParser(body); err != nil {
		return c.Status(http.StatusBadRequest).JSON("invalid json")
	}

	err := db.UpdateByID("users", c.Params("id"), body, &result)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.JSON(result)
}
