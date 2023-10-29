package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/alpaca-techwave/alpaca-mall-backend/middlewares"
)

func ProductRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	router.Post("/product/create", jwt, controllers.CreateProduct)
	router.Get("/product/index", controllers.GetAllProduct)
	router.Get("/product/find", controllers.GetBySearch)
}
