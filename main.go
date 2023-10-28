package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/swaggo/fiber-swagger"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	_ "github.com/alpaca-techwave/alpaca-mall-backend/docs"
	"github.com/alpaca-techwave/alpaca-mall-backend/routes"
)

// @title Alpaca Mall
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @host localhost:1323/api
// @BasePath /
func main() {
	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	config.Connect()

	routes.AuthRoutes(api)
	routes.UserRoutes(api)
	routes.AddressRoutes(api)

	err := app.Listen(":1323")
	if err != nil {
		log.Fatalf("fiber.Listen failed %s", err)
	}
}
