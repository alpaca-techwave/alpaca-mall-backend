package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	OrderID  uuid.UUID `                                json:"order_id"`
	SkuID    uuid.UUID `                                json:"sku_id"`
	Quantity uint32    `                                json:"quantity"`
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	oi.ID = uuid.New()
	return
}
