package controllers

import (
	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/AdrianAdame/imbedla-backend-fiber/platform/database"
	"github.com/gofiber/fiber/v2"
)

func getAllRoomsByUser (c *fiber.Ctx) error {
	user := &models.RoomUser{}

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	foundedRooms, err := db.getRoomsByUserId(user.ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error" : false,
		"data" : foundedRooms,
	})
}