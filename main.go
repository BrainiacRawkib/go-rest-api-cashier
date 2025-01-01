package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Instantiate Fiber
	app := fiber.New()

	// http handler
	app.Get("/testAPi", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Fiber first application!",
		})
	})

	// Run server
	app.Listen(":4433")
}
