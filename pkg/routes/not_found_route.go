package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"eror": true,
				"mgs":  "sorry, endpoint is not found",
			})
		},
	)
}