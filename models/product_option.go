package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductOption struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"   json:"id"`
	ProductID uuid.UUID `                                  json:"product_id"`
	Name      string    `                                  json:"name"`
	Sku       Sku       `gorm:"foreignKey:ProductOptionID"`
}

func (po *ProductOption) BeforeCreate(tx *gorm.DB) (err error) {
	po.ID = uuid.New()
	return
}
