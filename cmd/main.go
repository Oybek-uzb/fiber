package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"strconv"
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

	// Match request starting with /api
	app.Use("/api", func(c *fiber.Ctx) error {
		_ = c.SendString("middleware")

		return c.Next()
	})

	app.Get("/api/:id", func(c *fiber.Ctx) error {
		fmt.Println(c.Params("id"))
		return nil
	})

	// Attach multiple handlers
	app.Use("/api", func(c *fiber.Ctx) error {
		c.Set("X-Custom-Header", strconv.Itoa(rand.Int()))
		return c.Next()
	}, func(c *fiber.Ctx) error {
		return c.Next()
	})

	app.Listen(":3000")
}
