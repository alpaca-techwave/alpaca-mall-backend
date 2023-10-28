package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID      uuid.UUID `                                json:"user_id"`
	ProductID   uuid.UUID `                                json:"product_id"`
	Score       uint8     `                                json:"score"`
	Description string    `                                json:"description"`
}
