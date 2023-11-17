package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sku struct {
	gorm.Model
	ID              uuid.UUID   `gorm:"type:char(36);primaryKey" json:"id"`
	ProductID       uuid.UUID   `                                json:"product_id"`
	ProductOptionID uuid.UUID   `                                json:"product_option_id"`
	Price           float64     `                                json:"price"`
	Quantity        uint32      `                                json:"quantity"`
	CartItems       []CartItem  `gorm:"foreignKey:SkuID"`
	OrderItems      []OrderItem `gorm:"foreignKey:SkuID"`
}

func (sku *Sku) BeforeCreate(tx *gorm.DB) (err error) {
	sku.ID = uuid.New()
	return
}
