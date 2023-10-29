package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
)

func AuthRoutes(router fiber.Router) {
	router.Post("/auth/login", controllers.Login)
	router.Post("/auth/admin-login", controllers.LoginAdmin)
}
