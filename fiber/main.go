package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello FROM GO FIBER")
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "About page",
		})
	})

	app.Get("/about/:name", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
		}
		return c.SendString("Which name?")
	})

	app.Get("/api/:name/:age/:country", func(c *fiber.Ctx) error {
		return c.SendString(c.Params("name") + " is " + c.Params("age") + " years old and is from " + c.Params("country") + " .")
	})

	app.Listen(":3000")
}
