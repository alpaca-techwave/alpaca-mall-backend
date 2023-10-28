package routes

import (
	"github.com/alpaca-techwave/alpaca-mall-backend/controllers"
	"github.com/alpaca-techwave/alpaca-mall-backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router) {
	jwt := middlewares.NewAuthMiddleware()
	router.Post("/user/create", controllers.CreateUser)
	router.Get("/user/info", jwt, controllers.GetUserInfo)
	router.Put("/user/reset-password", jwt, controllers.UpdatePassword)
	router.Put("/user/update-info", jwt, controllers.UpdateUserInfo)
}
