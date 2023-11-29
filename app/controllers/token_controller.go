package controllers

import (
	"time"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/utils"
	"github.com/AdrianAdame/imbedla-backend-fiber/platform/database"
	"github.com/gofiber/fiber/v2"
)

func RenewTokens(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expiresAccesToken := claims.Expires

	if now > expiresAccesToken {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	renew := &models.Renew{}

	if err := c.BodyParser(renew); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expiresRefreshToken, err := utils.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if now < expiresRefreshToken {
		userID := claims.UserID

		db, err := database.OpenDBConnection()

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		foundedUser, err := db.GetUserById(userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		credentials, err := utils.GetCredentialsByRole(foundedUser.UserRole)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		tokens, err := utils.GenerateNewToken(userID.String(), credentials)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"tokens": fiber.Map{
				"access":  tokens.Access,
				"refresh": tokens.Refresh,
			},
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, your session was ended earlier",
		})
	}
}
