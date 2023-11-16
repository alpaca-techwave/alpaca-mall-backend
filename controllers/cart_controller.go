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

	var cartResponse models.GetCartResponse
	for _, cartItem := range cart.CartItems {
		sku := models.Sku{ID: cartItem.SkuID}
		result = config.Database.Find(&sku)
		if result.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "Not found Sku",
			})
		}
		product := models.Product{ID: sku.ProductID}
		result = config.Database.Find(&product)
		if result.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "Not found Product",
			})
		}
		productOption := models.ProductOption{ID: sku.ProductOptionID}
		result = config.Database.Find(&productOption)
		if result.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "Not found Product Option",
			})
		}
		cartItemResponse := models.GetCartItemResponse{
			CartItemID:        cartItem.ID,
			Image:             product.Image,
			ProductName:       product.Name,
			ProductOptionName: productOption.Name,
			Price:             uint32(sku.Price),
			Quantity:          cartItem.Quantity,
		}
		cartResponse.GetCartItemResposes = append(
			cartResponse.GetCartItemResposes,
			cartItemResponse,
		)
	}

	return c.Status(200).JSON(&cartResponse)
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
	result := config.Database.Where(models.Cart{UserID: userId}).Find(&cart)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Cart",
		})
	}

	var sku models.Sku
	result = config.Database.Find(&sku, cartItemBody.SkuID)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found Item",
		})
	}

	var cartItemExist models.CartItem
	result = config.Database.Where(models.CartItem{SkuID: cartItemBody.SkuID, CartID: cart.ID}).
		Find(&cartItemExist)
	if result.RowsAffected > 0 {
		cartItemExist.Quantity += cartItemBody.Quantity
		config.Database.Save(&cartItemExist)
		return c.Status(200).JSON(&cartItemExist)
	}

	cartItem := models.CartItem{
		ID:       uuid.New(),
		CartID:   cart.ID,
		SkuID:    cartItemBody.SkuID,
		Quantity: cartItemBody.Quantity,
	}
	config.Database.Create(&cartItem)
	return c.Status(200).JSON(&cartItem)
}

// @description Set Cart Item
// @tags        Cart
// @accept     json
// @produce    json
// @param cartitemid path string true "cart item id"
// @param cartitem body models.UpdateCartItemRequest true "cart item object"
// @success    200  {object}  models.CartItem
// @failure    500
// @failure    503
// @router     /cart/set/{cartItemId} [put]
// @security   BearerAuth
func SetCartItem(c *fiber.Ctx) error {
	cartItemIdStr := c.Params("cartItemId")
	cartItemId, err := uuid.Parse(cartItemIdStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Cart Item ID",
		})
	}

	cartItemBody := new(models.UpdateCartItemRequest)
	if err := c.BodyParser(cartItemBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}

	var cartItem models.CartItem
	result := config.Database.Find(&cartItem, cartItemId)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Cart Item",
		})
	}
	cartItem.Quantity = cartItemBody.Quantity
	config.Database.Save(&cartItem)
	return c.Status(200).JSON(&cartItem)
}

// @description Remove Cart Item
// @tags        Cart
// @accept     json
// @produce    json
// @param cartItemId path string true "cart item id"
// @success    200
// @failure    500
// @failure    503
// @router     /cart/remove/{cartItemId} [delete]
// @security   BearerAuth
func RemoveCartItem(c *fiber.Ctx) error {
	cartItemIdStr := c.Params("cartItemId")
	cartItemId, err := uuid.Parse(cartItemIdStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Cart Item ID",
		})
	}

	var cartItem models.CartItem
	result := config.Database.Find(&cartItem, cartItemId)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found Cart Item",
		})
	}
	config.Database.Delete(&cartItem)
	return c.Status(200).JSON(fiber.Map{
		"message": "Delete OK",
	})
}
