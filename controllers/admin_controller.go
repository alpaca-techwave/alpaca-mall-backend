package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

// @Description Create admin
// @Tags       Admin
// @Accept     json
// @Produce    json
// @Param admin body models.CreateAdminRequest true "Admin object"
// @Success    200  {object}  models.Admin
// @Failure    500
// @Failure    503
// @Router     /admin/create [post]
func CreateAdmin(c *fiber.Ctx) error {
	adminBody := new(models.CreateAdminRequest)
	if err := c.BodyParser(adminBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}
	godotenv.Load()
	if adminBody.SecretKey != os.Getenv("ADMIN_SECRET") {
		return c.Status(500).JSON(fiber.Map{
			"message": "No Permission",
		})
	}
	admin := models.Admin{
		ID:    uuid.New(),
		Email: adminBody.Email,
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(adminBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invaid Password Hashing",
		})
	}
	admin.Password = string(bytes)
	config.Database.Create(&admin)
	return c.Status(200).JSON(&admin)
}
