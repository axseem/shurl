package main

import (
	"log"
	"os"

	"github.com/axseem/shurl/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	app := fiber.New()
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Wellcome to shurl! POST to \"/shurl\" endpoint with {\"url\": \"YOUR URL HERE\"} in body")
	})
	app.Post("/", controllers.CreateShurl)
	app.Get("/:id", controllers.Redirect)

	app.Listen(os.Getenv("PORT"))
}
