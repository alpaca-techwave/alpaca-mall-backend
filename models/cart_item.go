package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	CartID   uuid.UUID `                                json:"cart_id"`
	SkuID    uuid.UUID `                                json:"sku_id"`
	Quantity uint32    `                                json:"quantity"`
}

type CreateCartItemRequest struct {
	SkuID    uuid.UUID `json:"sku_id"`
	Quantity uint32    `json:"quantity"`
}

type UpdateCartItemRequest struct {
	Quantity uint32 `json:"quantity"`
}

type GetCartItemResponse struct {
	CartItemID        uuid.UUID
	Image             string
	ProductName       string
	ProductOptionName string
	Price             uint32
	Quantity          uint32
}
