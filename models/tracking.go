package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tracking struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `                                json:"user_id"`
	ProductID uuid.UUID `                                json:"product_id"`
	Status    string    `                                json:"status"`
	Payment   string    `                                json:"payment"`
}
