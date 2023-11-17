package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/alpaca-techwave/alpaca-mall-backend/middlewares"
)

func OrderRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	router.Post("/order/create", jwt, controllers.CreateOrder)
	router.Get("/order", jwt, controllers.GetOrders)
}
