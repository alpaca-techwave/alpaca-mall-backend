package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

// @Description Create user
// @Tags       User
// @Accept     json
// @Produce    json
// @Param user body models.CreateUserRequest true "User object"
// @Success    200  {object}  models.User
// @Failure    500
// @Failure    503
// @Router     /user/create [post]
func CreateUser(c *fiber.Ctx) error {
	userBody := new(models.CreateUserRequest)
	if err := c.BodyParser(userBody); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	user := models.User{
		ID:        uuid.New(),
		Email:     userBody.Email,
		FirstName: userBody.FirstName,
		LastName:  userBody.LastName,
		Tel:       userBody.Tel,
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invaid Password Hashing",
		})
	}
	user.Password = string(bytes)
	config.Database.Create(&user)
	return c.Status(200).JSON(user)
}

// @Description Get user by ID
// @Tags       User
// @Accept     json
// @Produce    json
// @Success    200  {object}  models.User
// @Failure    500
// @Failure    503
// @Router     /user/info [get]
// @Security   BearerAuth
func GetUserInfo(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)
	idStr, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	var user models.User
	result := config.Database.Preload("Address").Find(&user, idStr)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found User",
		})
	}
	return c.Status(200).JSON(&user)
}

// @Description Reset password
// @Tags       User
// @Accept     json
// @Produce    json
// @Param user body models.ResetPasswordRequest true "Reset Password object"
// @Success    200
// @Failure    500
// @Failure    503
// @Router     /user/reset-password [put]
// @Security   BearerAuth
func UpdatePassword(c *fiber.Ctx) error {
	resetPassword := new(models.ResetPasswordRequest)
	if err := c.BodyParser(resetPassword); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}

	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)
	idStr, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var user models.User
	result := config.Database.Find(&user, idStr)
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not found User",
		})
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(resetPassword.CurrentPassword),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Password",
		})
	}

	bytes, err := bcrypt.GenerateFromPassword(
		[]byte(resetPassword.NewPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invaid Password Hashing",
		})
	}

	config.Database.Model(&user).Updates(models.User{Password: string(bytes)})
	return c.Status(200).JSON(fiber.Map{
		"message": "Successful reset password",
	})
}

// @Description Update user info
// @Tags       User
// @Accept     json
// @Produce    json
// @Param user body models.UpdateInfoRequest true "Update info object"
// @Success    200
// @Failure    500
// @Failure    503
// @Router     /user/update-info [put]
// @Security   BearerAuth
func UpdateUserInfo(c *fiber.Ctx) error {
	updateInfo := new(models.UpdateInfoRequest)
	if err := c.BodyParser(updateInfo); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}

	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)
	idStr, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var user models.User
	result := config.Database.Find(&user, idStr)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found User",
		})
	}
	config.Database.Model(&user).
		Updates(models.User{Email: updateInfo.Email, FirstName: updateInfo.FirstName, LastName: updateInfo.LastName, Tel: updateInfo.Tel})
	return c.Status(200).JSON(&user)
}
