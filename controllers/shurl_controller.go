package controllers

import (
	"os"

	"github.com/axseem/shurl/database/models"
	"github.com/axseem/shurl/database/queries"
	"github.com/axseem/shurl/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jaevor/go-nanoid"
)

func CreateShurl(c *fiber.Ctx) error {
	body := &models.Shurl{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can't parse JSON"})
	}

	utils.PrefixHTTP(&body.Url)

	if !utils.IsValidURL(body.Url) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid URL"})
	}

	var id string
	if body.Shurl == "" {
		if body.Length == 0 {
			body.Length = 5
		}
		generateID, err := nanoid.Standard(body.Length)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}
		id = generateID()
	} else {
		if !utils.IsValidID(body.Shurl) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "shurl can't contain other characters than: " + utils.ALLOWED_CHARS})
		}
		id = body.Shurl
		body.Length = len(body.Shurl)
	}

	shurl, err := queries.GetShurl(id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
	if shurl != "" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "shurl is already in use. Try again"})
	}

	if err = queries.AddShurl(id, body.Url, body.Expiry); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to add shurl to db"})
	}

	return c.Status(fiber.StatusOK).JSON(
		models.Shurl{
			Url:    body.Url,
			Shurl:  os.Getenv("DOMAIN") + "/" + id,
			Length: body.Length,
			Expiry: body.Expiry,
		},
	)
}
