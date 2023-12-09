package routes

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/controllers"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	// Create Routes Group
	route := a.Group("/api")

	/** User Authentication Routes */
	route.Post("/user/logout", middleware.JWTProtected(), controllers.UserSignOut) // de-authorization user
	route.Post("/token/refresh", middleware.JWTProtected(), controllers.RenewTokens)   // renew Access & Refresh tokens

	/** User Rooms Routes */
	route.Get("/rooms/user/:userId", middleware.JWTProtected(), controllers.GetAllRoomsByUser)
	route.Get("/rooms/:id", middleware.JWTProtected(), controllers.GetRoomById)
	route.Post("/rooms", middleware.JWTProtected(), controllers.CreateNewRoomByUserId)
	route.Patch("/rooms", middleware.JWTProtected(), controllers.UpdateRoomById)
	route.Delete("/rooms", middleware.JWTProtected(), controllers.DeleteRoomById)

	/** User plants Routes */
	route.Get("/plants/room/:roomId", middleware.JWTProtected(), controllers.GetAllPlantsByRoom)
	route.Get("/plants/:id", middleware.JWTProtected(), controllers.GetPlantById)
	route.Post("/plants", middleware.JWTProtected(), controllers.CreateNewPlant)
	route.Patch("/plants", middleware.JWTProtected(), controllers.UpdatePlantById)
	route.Delete("/plants", middleware.JWTProtected(), controllers.DeletePlantById)
}