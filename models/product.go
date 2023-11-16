package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID             uuid.UUID       `gorm:"type:char(36);primaryKey" json:"id"`
	Name           string          `                                json:"name"`
	Description    string          `                                json:"description"`
	Image          string          `                                json:"img"`
	Trackings      []Tracking      `gorm:"foreignKey:ProductID"`
	Reviews        []Review        `gorm:"foreignKey:ProductID"`
	ProductOptions []ProductOption `gorm:"foreignKey:ProductID"`
	Skus           []Sku           `gorm:"foreignKey:ProductID"`
}

type CreateProductRequest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Image          string `json:"img"`
	ProductOptions []struct {
		Name string `json:"name"`
		Skus struct {
			Price    float64 `                                json:"price"`
			Quantity uint32  `                                json:"quantity"`
		}
	}
}
