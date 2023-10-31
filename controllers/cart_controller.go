package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

// @Description Get user Cart
// @Tags       Cart
// @Accept     json
// @Produce    json
// @Success    200  {array}  models.Cart
// @Failure    500
// @Failure    503
// @Router     /cart [get]
// @Security   BearerAuth
func GetMyCart(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	userIdStr := claims["ID"].(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	var cart models.Cart
	result := config.Database.Preload("CartItems").Where(models.Cart{UserID: userId}).Find(&cart)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Cart",
		})
	}
	return c.Status(200).JSON(&cart)
}

// @Description Add Cart Item
// @Tags        Cart
// @Accept     json
// @Produce    json
// @Param cartItem body models.CreateCartItemRequest true "Cart Item object"
// @Success    200  {object}  models.CartItem
// @Failure    500
// @Failure    503
// @Router     /cart/add [post]
// @Security   BearerAuth
func AddCartItem(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	userIdStr := claims["ID"].(string)
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}
	cartItemBody := new(models.CreateCartItemRequest)
	if err := c.BodyParser(cartItemBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}

	var cart models.Cart
	result := config.Database.Preload("CartItems").Where(models.Cart{UserID: userId}).Find(&cart)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Cart",
		})
	}
	cart.Quantity += cartItemBody.Quantity
	config.Database.Save(&cart)

	var existCartItem models.CartItem
	result = config.Database.Where(models.CartItem{
		ProductID:   cartItemBody.ProductID,
		VariantName: cartItemBody.VariantName, SkuName: cartItemBody.SkuName,
	}).Find(&existCartItem)
	if result.RowsAffected != 0 {
		existCartItem.Quantity += cartItemBody.Quantity
		config.Database.Save(&existCartItem)
		return c.Status(200).JSON(&existCartItem)
	}

	cartItem := models.CartItem{
		ID:          uuid.New(),
		CartID:      cart.ID,
		ProductID:   cartItemBody.ProductID,
		VariantName: cartItemBody.VariantName,
		SkuName:     cartItemBody.SkuName,
		Quantity:    cartItemBody.Quantity,
	}
	config.Database.Create(&cartItem)
	return c.Status(200).JSON(&cartItem)
}
