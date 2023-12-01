package controllers

import (
	"time"

	"github.com/AdrianAdame/imbedla-backend-fiber/app/models"
	"github.com/AdrianAdame/imbedla-backend-fiber/pkg/utils"
	"github.com/AdrianAdame/imbedla-backend-fiber/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserSignUp(c *fiber.Ctx) error {
	// Create a new user auth struct
	signUp := &models.SignUp{}

	signUp.Role = "user"

	// Checking received data from JSON body.
	if err := c.BodyParser(signUp); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate sign up fields
	if err := validate.Struct(signUp); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking role from sign up data.
	role, err := utils.VerifyRole(signUp.Role)

	if err != nil {
		// Return status 400 and message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	user := &models.User{}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Email = signUp.Email
	user.Firstname = signUp.Firstname
	user.Lastname = signUp.Lastname
	user.About = signUp.About
	user.PasswordHash = utils.GeneratePassword(signUp.Password)
	user.UserStatus = 1 // 0 == blocked, 1 == activate
	user.UserRole = role

	// Validate user fields
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorError(err),
		})
	}
	// Create a new user with validated data
	if err := db.CreateUser(user); err != nil {

		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Delete password hash field from JSON view.
	user.PasswordHash = ""

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "registered",
		"user":  user,
	})
}

func UserSignIn(c *fiber.Ctx) error {
	// Create a new user Auth struct
	signIn := &models.SignIn{}

	// Checking received data from JSON body
	if err := c.BodyParser(signIn); err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user by email
	foundedUser, err := db.GetUserByEmail(signIn.Email)
	if err != nil {
		// Return status 500 and database error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given email is not found",
		})
	}

	// compare given user password with stored in found user
	compareUserPassword := utils.ComparePasswords(foundedUser.PasswordHash, signIn.Password)
	if !compareUserPassword {
		// Return , if  passwordis not compare to stored in db
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "wrong user email address or password",
		})
	}

	// Get role credentials from founded user
	credentials, err := utils.GetCredentialsByRole(foundedUser.UserRole)
	if err != nil {
		// Return status 400 and error message
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewToken(foundedUser.ID.String(), credentials)
	if err != nil {
		// Return 500 and token generation error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	resToken := &models.Tokens{}

	resToken.TokenID = uuid.New()
	resToken.UserID = foundedUser.ID
	resToken.Access = tokens.Access
	resToken.Refresh = tokens.Refresh
	resToken.CreatedAt = time.Now()

	if err := db.CreateTokens(resToken); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  tokens.Access,
			"refresh": tokens.Refresh,
		},
	})
}

func UserSignOut(c *fiber.Ctx) error {

	// Get claims from JWT
	claims, err := utils.ExtractTokenMetadata(c)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token := &models.Tokens{}

	token.UserID = claims.UserID

	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := db.DeleteTokens(token); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
