package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
)

func AdminRoutes(router fiber.Router) {
	router.Post("/admin/create", controllers.CreateAdmin)
}
