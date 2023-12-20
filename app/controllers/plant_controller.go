package controllers

import (
	"encoding/json"
	"time"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/utils"
	"github.com/AdrianAdame/imbedla-backend-fiber/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllPlantsByRoom(c *fiber.Ctx) error {
	if err := c.Params("roomId"); err == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
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

	roomId, _ := uuid.Parse(c.Params("roomId"))

	foundedPlants, err := db.GetAllPlantsByRoomId(roomId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  foundedPlants,
	})
}

func GetPlantById(c *fiber.Ctx) error {
	if err := c.Params("id"); err == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
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

	plantId, _ := uuid.Parse(c.Params("id"))

	plantRow, err := db.GetPlantById(plantId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "plant with the given id was not found",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"room":  plantRow,
	})
}

func CreateNewPlant(c *fiber.Ctx) error {
	plantH := &models.PlantH{}

	if err := c.BodyParser(plantH); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(plantH); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}

	json.Unmarshal(plantH.ModuleInformation, &plantH.ModuleInformation)
	json.Unmarshal(plantH.ModuleSpecs, &plantH.ModuleSpecs)

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	record := &models.PlantD{}

	record.ID = uuid.New()
	record.UserId = plantH.UserId
	record.RoomId = plantH.RoomId
	record.Name = plantH.Name
	record.RefPlant = plantH.RefPlant
	record.CreatedAt = time.Now()
	record.UpdatedAt = time.Now()
	record.ModuleInformation = string(plantH.ModuleInformation)
	record.ModuleSpecs = string(plantH.ModuleSpecs)
	record.Favorite = false

	if err := validate.Struct(record); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}

	if err := db.CreatePlant(record); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "created",
		"plant": record,
	})
}

func UpdatePlantById(c *fiber.Ctx) error {
	plant := &models.UpdatePlantH{}

	if err := c.BodyParser(plant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(plant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	foundedPlant, err := db.GetPlantById(plant.ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if plant.Name != "" {
		foundedPlant.Name = plant.Name
	}

	if plant.ModuleInformation != nil {
		json.Unmarshal(plant.ModuleInformation, &plant.ModuleInformation)

		foundedPlant.ModuleInformation = string(plant.ModuleInformation)
	}

	if plant.ModuleSpecs != nil {
		json.Unmarshal(plant.ModuleSpecs, &plant.ModuleSpecs)

		foundedPlant.ModuleSpecs = string(plant.ModuleSpecs)
	}

	foundedPlant.UpdatedAt = time.Now()
	foundedPlant.Favorite = plant.Favorite

	if err := validate.Struct(foundedPlant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}

	if err := db.EditPlant(&foundedPlant); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "modified",
		"plant": foundedPlant,
	})
}

func DeletePlantById(c *fiber.Ctx) error {
	plantToDelete := &models.DeletePlant{}

	if err := c.BodyParser(plantToDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()

	if err := validate.Struct(plantToDelete); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}

	db, err := database.OpenDBConnection()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, err = db.GetPlantById(plantToDelete.ID)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.DeletePlant(plantToDelete.ID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "deleted",
	})
}

func GetAllPlantsByUser(c *fiber.Ctx) error {
	if err := c.Params("userId"); err == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
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

	userId, _ := uuid.Parse(c.Params("userId"))

	foundedPlants, err := db.GetFavoritePlantsByUserId(userId)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  foundedPlants,
	})
}
