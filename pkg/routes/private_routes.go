package routes

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/controllers"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	// Create Routes Group
	route := a.Group("/api")

	route.Get("/", middleware.JWTProtected(), func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens
}
