package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"

	"github.com/alpaca-techwave/alpaca-mall-backend/models"
	"github.com/alpaca-techwave/alpaca-mall-backend/repository"
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

	user, err := repository.FindByUserCredentials(loginRequest.Email, loginRequest.Password)
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

// @Description Login admin
// @Tags       Auth
// @Accept     json
// @Produce    json
// @Param user body models.LoginRequest true "Token"
// @Success    200  {object}  models.LoginResponse
// @Failure    500
// @Failure    503
// @Router     /auth/admin-login [post]
func LoginAdmin(c *fiber.Ctx) error {
	godotenv.Load()
	loginRequest := new(models.LoginRequest)
	if err := c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	admin, err := repository.FindByAdminCredentials(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	day := time.Hour * 24
	claims := jtoken.MapClaims{
		"ID":    admin.ID,
		"email": admin.Email,
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
