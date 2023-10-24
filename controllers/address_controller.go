package controllers

import (
	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

//  @Description Create address
//  @Tags       Address
//  @Accept     json
//  @Produce    json
//  @Param user body models.Address true "Address object"
//  @Success    200  {object}  models.Address
//  @Failure    500
//  @Failure    503
//  @Router     /address/create [post]
func CreateAddress(c *fiber.Ctx) error {
  address := new(models.Address)
  address.ID = uuid.New()
  if err := c.BodyParser(address); err != nil {
    return c.Status(503).SendString(err.Error())
  }
  config.Database.Create(&address)
  return c.Status(200).JSON(address)

}
