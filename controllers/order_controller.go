package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

// @Description Create Order
// @Tags       Order
// @Accept     json
// @Produce    json
// @Param address body models.CreateOrderRequest true "Order object"
// @Success    200  {object}  models.Order
// @Failure    500
// @Failure    503
// @Router     /order/create [post]
// @Security   BearerAuth
func CreateOrder(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	userIdStr := claims["ID"].(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	orderBody := new(models.CreateOrderRequest)
	if err := c.BodyParser(&orderBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}

	var cart models.Cart
	result := config.Database.Preload("CartItems").Where(models.Cart{UserID: userId}).Find(&cart)
	if result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Not Found Cart",
		})
	}

	order := models.Order{
		ID:      uuid.New(),
		Status:  "WAIT",
		UserID:  userId,
		Payment: orderBody.Payment,
	}
	config.Database.Create(&order)

	for _, item := range cart.CartItems {
		orderItem := models.OrderItem{
			OrderID:  order.ID,
			SkuID:    item.SkuID,
			Quantity: item.Quantity,
		}
		config.Database.Create(&orderItem)
		config.Database.Delete(&item)
	}

	return c.Status(200).JSON(&order)
}

// @Description Get All Orders
// @Tags       Order
// @Accept     json
// @Produce    json
// @Success    200  {array}  models.Order
// @Failure    500
// @Failure    503
// @Router     /order [get]
// @Security   BearerAuth
func GetOrders(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	userIdStr := claims["ID"].(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var orders []models.Order
	config.Database.Preload("OrderItems").Where(models.Order{UserID: userId}).Find(&orders)
	return c.Status(200).JSON(&orders)
}
