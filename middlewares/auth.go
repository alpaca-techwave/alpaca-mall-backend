package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

// Middleware JWT function
func NewAuthMiddleware() fiber.Handler {
  godotenv.Load()
 return jwtware.New(jwtware.Config{
  SigningKey: []byte(os.Getenv("JWT_SECRET")),
 })
}
