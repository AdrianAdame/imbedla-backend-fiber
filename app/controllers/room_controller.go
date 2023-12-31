package controllers

import (
	"time"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/utils"
	"github.com/AdrianAdame/imbedla-backend-fiber/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllRoomsByUser (c *fiber.Ctx) error {
	if err := c.Params("userId"); err == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : err,
		})
	}

	userId, _ := uuid.Parse(c.Params("userId"))

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	foundedRooms, err := db.GetRoomsByUserId(userId)

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

func GetRoomById (c *fiber.Ctx) error {
	if err := c.Params("id"); err == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : err,
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		// Return status 500 and database error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	roomId, _ := uuid.Parse(c.Params("id"))

	room, err := db.GetRoomById(roomId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "room with the given id was not found",
		})
	}

	return c.JSON(fiber.Map{
		"error" : false,
		"room" : room,
	})
}

func CreateNewRoomByUserId ( c *fiber.Ctx) error {
	room := &models.Room{}

	if err := c.BodyParser(room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : utils.ValidatorError(err),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	roomTemp := &models.Room{}

	roomTemp.ID = uuid.New()
	roomTemp.CreatedAt = time.Now()
	roomTemp.UpdatedAt = time.Now()
	roomTemp.Name = room.Name
	roomTemp.Color = room.Color
	roomTemp.Type = room.Type
	roomTemp.UserId = room.UserId

	if err := validate.Struct(roomTemp); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : utils.ValidatorError(err),
		})
	}

	if err := db.CreateRoom(roomTemp); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error" : false,
		"msg" : "created",
		"room" : roomTemp,
	})
}

func UpdateRoomById (c *fiber.Ctx) error {
	room := &models.UpdateRoom{}

	if err := c.BodyParser(room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(room); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : utils.ValidatorError(err),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	foundedRoom, err := db.GetRoomById(room.ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	if room.Name != "" {
		foundedRoom.Name = room.Name
	}

	if room.Color != "" {
		foundedRoom.Color = room.Color
	}

	if room.Type != "" {
		foundedRoom.Type = room.Type
	}

	foundedRoom.UpdatedAt = time.Now()

	if err := validate.Struct(foundedRoom); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : utils.ValidatorError(err),
		})
	}

	if err := db.EditRoom(&foundedRoom); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error" : false,
		"msg" : "modified",
		"room" : foundedRoom,
	})
}

func DeleteRoomById (c *fiber.Ctx) error {
	roomToDelete := &models.DeleteRoom{}

	if err := c.BodyParser(roomToDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(roomToDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : true,
			"msg" : utils.ValidatorError(err),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	_, err = db.GetRoomById(roomToDelete.ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	if err := db.DeleteRoom(roomToDelete.ID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error" : true,
			"msg" : err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error" : false,
		"msg" : "deleted",
	})
}