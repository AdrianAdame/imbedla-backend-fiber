package utils

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenMetadata struct {
	UserID      uuid.UUID
	Credentials map[string]bool
	Expires     int64
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		expires := int64(claims["expires"].(float64))

		credencials := map[string]bool{}

		return &TokenMetadata{
			UserID:      userID,
			Credentials: credencials,
			Expires:     expires,
		}, nil
	}
	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	//Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, fiber.ErrBadGateway
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET_KEY")), nil
}