package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID         uuid.UUID   `gorm:"type:char(36);primaryKey" json:"id"`
	UserID     uuid.UUID   `                                json:"user_id"`
	Status     string      `                                json:"status"`
	Payment    string      `                                json:"payment"`
	OrderItems []OrderItem `gorm:"foreignkey:OrderID"`
}

type CreateOrderRequest struct {
	Payment string `json:"payment"`
}
