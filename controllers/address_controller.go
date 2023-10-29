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
// @Param address body models.AddressRequest true "Address object"
// @Success    200  {object}  models.Address
// @Failure    500
// @Failure    503
// @Router     /address/create [post]
// @Security   BearerAuth
func CreateAddress(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	idStr := claims["ID"].(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	addressBody := new(models.AddressRequest)
	if err := c.BodyParser(addressBody); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}
	if addressBody.IsDefault == true {
		var defaultAddress models.Address
		result := config.Database.Where(models.Address{UserID: id, IsDefault: true}).
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
		UserID:      id,
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
	idStr := claims["ID"].(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	var address models.Address
	result := config.Database.Where(models.Address{UserID: id, IsDefault: true}).Find(&address)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Default Address",
		})
	}
	return c.Status(200).JSON(&address)
}

// @Description Create address
// @Tags       Address
// @Accept     json
// @Produce    json
// @Param id path string true "Address ID"
// @Param address body models.AddressRequest true "Address object"
// @Success    200  {object}  models.Address
// @Failure    500
// @Failure    503
// @Router     /address/update/{id} [put]
// @Security   BearerAuth
func UpdateAddress(c *fiber.Ctx) error {
	addressBody := new(models.AddressRequest)
	if err := c.BodyParser(addressBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}

	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	userIdStr := claims["ID"].(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid UserID",
		})
	}
	if addressBody.IsDefault == true {
		var defaultAddress models.Address
		result := config.Database.Where(models.Address{UserID: userId, IsDefault: true}).
			Find(&defaultAddress)
		if result.RowsAffected != 0 {
			defaultAddress.IsDefault = false
			config.Database.Save(&defaultAddress)
		}
	}

	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	var updateAddress models.Address
	result := config.Database.Find(&updateAddress, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Address",
		})
	}
	config.Database.Model(&updateAddress).
		Updates(models.Address{Name: addressBody.Name, Tel: addressBody.Tel, MainAddress: addressBody.MainAddress, SubAddress: addressBody.SubAddress, IsDefault: addressBody.IsDefault})
	return c.Status(200).JSON(&updateAddress)
}

// @Description Remove address
// @Tags       Address
// @Accept     json
// @Produce    json
// @Param id path string true "Address ID"
// @Success    200
// @Failure    500
// @Failure    503
// @Router     /address/remove/{id} [delete]
// @Security   BearerAuth
func RemoveAddress(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	var address models.Address
	result := config.Database.Find(&address, id)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Address",
		})
	}
	if address.IsDefault == true {
		return c.Status(500).JSON(fiber.Map{
			"message": "Can not remove Default Address",
		})
	}
	config.Database.Delete(&address)
	return c.Status(200).JSON(fiber.Map{
		"message": "Successful Remove Address",
	})
}
