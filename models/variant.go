package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	ProductID uuid.UUID `                                json:"product_id"`
	Name      string    `                                json:"name"`
	Quantity  uint8     `                                json:"quantity"`
	Skus      []Sku     `gorm:"foreignKey:VariantID"`
}

func (variant *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	variant.ID = uuid.New()
	return
}
