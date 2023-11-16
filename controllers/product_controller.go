package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

// @Description Create user
// @Tags       Product
// @Accept     json
// @Produce    json
// @Param product body models.CreateProductRequest true "Product object"
// @Success    200 {object}  models.Product
// @Failure    500
// @Failure    503
// @Router     /product/create [post]
// @Security   BearerAuth
func CreateProduct(c *fiber.Ctx) error {
	userToken := c.Locals("user").(*jtoken.Token)
	claims := userToken.Claims.(jtoken.MapClaims)
	idStr := claims["ID"].(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid ID",
		})
	}

	var admin models.Admin
	result := config.Database.Find(&admin, id)
	if result.RowsAffected == 0 {
		return c.Status(500).JSON(fiber.Map{
			"message": "No Permission",
		})
	}

	productBody := new(models.CreateProductRequest)
	if err := c.BodyParser(productBody); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Body",
		})
	}
	product := models.Product{
		ID:          uuid.New(),
		Name:        productBody.Name,
		Description: productBody.Description,
		Image:       productBody.Image,
	}
	config.Database.Create(&product)
	for _, option := range productBody.ProductOptions {
		productOption := models.ProductOption{
			ID:        uuid.New(),
			ProductID: product.ID,
			Name:      option.Name,
		}
		config.Database.Create(&productOption)
		skus := models.Sku{
			ID:              uuid.New(),
			ProductID:       product.ID,
			ProductOptionID: productOption.ID,
			Price:           option.Skus.Price,
			Quantity:        option.Skus.Quantity,
		}
		config.Database.Create(&skus)
	}
	return c.Status(200).JSON(&product)
}

// @Description Get All Product
// @Tags       Product
// @Accept     json
// @Produce    json
// @Success    200  {array}  models.Product
// @Failure    500
// @Failure    503
// @Router     /product/index [get]
func GetAllProduct(c *fiber.Ctx) error {
	var products []models.Product
	config.Database.Preload("ProductOptions.Sku").
		Preload("Skus").
		Order("created_at asc").
		Find(&products)
	return c.JSON(&products)
}

// @Description Get Product by Search
// @Tags       Product
// @Accept     json
// @Produce    json
// @Param search query string true "Search term"
// @Success    200  {array}  models.Product
// @Failure    500
// @Failure    503
// @Router     /product/find [get]
func GetBySearch(c *fiber.Ctx) error {
	searchTerm := c.Query("search")
	searchTerm = "%" + searchTerm + "%"
	var products []models.Product
	config.Database.Preload("ProductOptions.Sku").
		Preload("Skus").
		Where("name LIKE ?", searchTerm).
		Find(&products)
	return c.JSON(&products)
}

// @Description Get Product By ID
// @Tags       Product
// @Accept     json
// @Produce    json
// @Param id path string true "Product ID"
// @Success    200 {object} models.Product
// @Failure    500
// @Failure    503
// @Router     /product/{id} [get]
func GetProductByID(c *fiber.Ctx) error {
	productIdStr := c.Params("id")
	productId, err := uuid.Parse(productIdStr)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Invalid Product ID",
		})
	}

	var product models.Product
	result := config.Database.Preload("ProductOptions.Sku").
		Preload("Skus").
		Find(&product, productId)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found Product",
		})
	}
	return c.JSON(&product)
}
