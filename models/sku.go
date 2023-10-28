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
	Value     uint8     `                                json:"value"`
}
