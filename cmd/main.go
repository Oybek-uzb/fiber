package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Static("/", "./internal")

	app.Get("/:param", func(c *fiber.Ctx) error {
		return c.SendString("param: " + c.Params("param"))
	})

	app.Use("/api", func(c *fiber.Ctx) error {
		_ = c.SendString("middleware")

		return c.Next()
	})

	app.Get("/api/:id", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("id"))
		return nil
	})

	app.Listen(":3000")
}
