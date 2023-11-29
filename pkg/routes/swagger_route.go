package routes

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")

	// Routes for GET method:
	route.Get("*", swagger.HandlerDefault)
}
