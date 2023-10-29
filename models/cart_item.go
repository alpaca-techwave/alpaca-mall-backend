package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	CartID      uuid.UUID `                                json:"cart_id"`
	ProductID   uuid.UUID `                                json:"product_id"`
	VariantName string    `                                json:"variant_name"`
	SkuName     string    `                                json:"sku_name"`
	Quantity    uint8     `                                json:"quantity"`
}
