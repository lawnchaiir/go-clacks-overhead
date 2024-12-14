package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-Clacks-Overhead", "GNU Terry Pratchett")
		return c.Next()
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
	})

	app.Listen(":8080")
}
