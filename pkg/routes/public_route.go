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

	// Temp

	route.Post("/auth/logout", controllers.UserSignOut) // de-authorization user

	// Routes for POST method:
	route.Post("/auth/register", controllers.UserSignUp) // register a new user
	route.Post("/auth/sign", controllers.UserSignIn)     // auth, return Access & Refresh tokens

}
