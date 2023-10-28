package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ProductID uuid.UUID `                                json:"product_id"`
	Color     string    `                                json:"color"`
	Skus      []Sku     `gorm:"foreignKey:VariantID"`
}
