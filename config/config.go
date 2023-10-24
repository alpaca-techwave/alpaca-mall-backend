package config

import (
	"os"

	"github.com/alpaca-techwave/alpaca-mall-backend/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() error {
  godotenv.Load()

  var err error

  Database, err = gorm.Open(mysql.Open(os.Getenv("DATABASE_URI")), &gorm.Config{
    SkipDefaultTransaction: true,
    PrepareStmt:            true,
  })

  if err != nil {
    panic(err)
  }

  Database.AutoMigrate(&models.User{},&models.Address{})

  return nil
}
