package routes

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api")

	route.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	// Routes for POST method:
	route.Post("/auth/register", controllers.UserSignUp) // register a new user
	route.Post("/auth/sign", controllers.UserSignIn)     // auth, return Access & Refresh tokens

	route.Get("/test1", func(c *fiber.Ctx) error {
		return c.SendString("is hot reloading")
	})
}
