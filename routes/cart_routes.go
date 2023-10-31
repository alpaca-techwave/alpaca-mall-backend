package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/alpaca-techwave/alpaca-mall-backend/middlewares"
)

func CartRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	router.Get("/cart", jwt, controllers.GetMyCart)
	router.Post("/cart/add", jwt, controllers.AddCartItem)
}
