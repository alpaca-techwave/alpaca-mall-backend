package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

// @Description Create address
// @Tags       Address
// @Accept     json
// @Produce    json
// @Param address body models.CreateAddressRequest true "Address object"
// @Success    200  {object}  models.Address
// @Failure    500
// @Failure    503
// @Router     /address/create [post]
// @Security   BearerAuth
func CreateAddress(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)
	idStr, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	addressBody := new(models.CreateAddressRequest)
	if err := c.BodyParser(addressBody); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	if addressBody.IsDefault == true {
		var defaultAddress models.Address
		result := config.Database.Where(models.Address{UserID: idStr, IsDefault: true}).
			Find(&defaultAddress)
		if result.RowsAffected != 0 {
			defaultAddress.IsDefault = false
			config.Database.Save(&defaultAddress)
		}
	}
	address := models.Address{
		ID:          uuid.New(),
		Name:        addressBody.Name,
		Tel:         addressBody.Tel,
		MainAddress: addressBody.MainAddress,
		SubAddress:  addressBody.SubAddress,
		IsDefault:   addressBody.IsDefault,
		UserID:      idStr,
	}
	config.Database.Create(&address)
	return c.Status(200).JSON(address)
}

// @Description Get default address
// @Tags       Address
// @Accept     json
// @Produce    json
// @Success    200  {object}  models.Address
// @Failure    500
// @Failure    503
// @Router     /address/default [get]
// @Security   BearerAuth
func GetDefaultAddress(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	id := claims["ID"].(string)
	idStr, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	var address models.Address
	result := config.Database.Where(models.Address{UserID: idStr, IsDefault: true}).Find(&address)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Default Address",
		})
	}
	return c.Status(200).JSON(&address)
}
