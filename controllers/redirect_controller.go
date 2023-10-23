package controllers

import (
	"github.com/axseem/shurl/database/queries"
	"github.com/gofiber/fiber/v2"
)

func Redirect(c *fiber.Ctx) error {
	id := c.Params("id")

	shurl, err := queries.GetShurl(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	if shurl == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "shurl doesn't exist"})
	}

	return c.Redirect(shurl)
}
