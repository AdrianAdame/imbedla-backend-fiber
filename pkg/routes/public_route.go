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
	route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens
}
