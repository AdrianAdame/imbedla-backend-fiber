package main

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/configs"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/middleware"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	// Middleware
	middleware.FiberMiddleware(app)

	// Endpoints route
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.SwaggerRoute(app)
	routes.NotFoundRoute(app)

	app.Listen(":3000")
}
