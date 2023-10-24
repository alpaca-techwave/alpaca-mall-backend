package routes

import (
	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/alpaca-techwave/alpaca-mall-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)


func AuthRoutes(router fiber.Router){
  jwt := middlewares.NewAuthMiddleware()
  router.Post("/auth/login",controllers.Login)
  router.Get("/auth/test",jwt,controllers.Protected)
}
