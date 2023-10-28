package controllers

import (
	"os"
	"time"

	"github.com/alpaca-techwave/alpaca-mall-backend/models"
	"github.com/alpaca-techwave/alpaca-mall-backend/repository"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// @Description Login
// @Tags       Auth
// @Accept     json
// @Produce    json
// @Param user body models.LoginRequest true "Token"
// @Success    200  {object}  models.LoginResponse
// @Failure    500
// @Failure    503
// @Router     /auth/login [post]
func Login(c *fiber.Ctx) error {
	godotenv.Load()
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(models.LoginResponse{
		Token: t,
	})
}

// @Description Test jwt auth
// @Tags       Auth
// @Accept     json
// @Produce    json
// @Success    200
// @Failure    500
// @Failure    503
// @Router     /auth/test [get]
// @Security   BearerAuth
func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	ID := claims["ID"].(string)
	return c.SendString("Welcome 👋" + email + " " + ID)
}
