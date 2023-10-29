package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sku struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	VariantID uuid.UUID `                                json:"variant_id"`
	Name      string    `                                json:"name"`
	Quantity  uint8     `                                json:"quantity"`
}

func (sku *Sku) BeforeCreate(tx *gorm.DB) (err error) {
	sku.ID = uuid.New()
	return
}
