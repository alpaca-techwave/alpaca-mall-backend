package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/alpaca-techwave/alpaca-mall-backend/middlewares"
)

func AddressRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	router.Post("/address/create", jwt, controllers.CreateAddress)
	router.Get("/address/default", jwt, controllers.GetDefaultAddress)
}
