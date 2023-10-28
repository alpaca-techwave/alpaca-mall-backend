package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Name        string     `                                json:"name"`
	Price       float32    `                                json:"price"`
	Description string     `                                json:"description"`
	Image       string     `                                json:"img"`
	Carts       []Cart     `gorm:"foreignKey:ProductID"`
	Variants    []Variant  `gorm:"foreignKey:ProductID"`
	Trackings   []Tracking `gorm:"foreignKey:ProductID"`
	Reviews     []Review   `gorm:"foreignKey:ProductID"`
}
