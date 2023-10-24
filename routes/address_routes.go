package routes

import (
	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(router fiber.Router){
  router.Post("/address/create",controllers.CreateAddress)
}
