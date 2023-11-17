package config

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/alpaca-techwave/alpaca-mall-backend/models"
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

	Database.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Address{},
		&models.Product{},
		&models.ProductOption{},
		&models.Sku{},
		&models.Cart{},
		&models.Review{},
		&models.Order{},
		&models.CartItem{},
		&models.OrderItem{},
	)

	return nil
}
