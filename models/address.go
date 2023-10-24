package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct{
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
  UserID uuid.UUID `json:"user_id"`
  Name string `json:"name"`
  Tel string `json:"tel"`
  MainAddress string `json:"main_address"`
  SubAddress string `json:"sub_address"`
  IsDefault bool `json:"is_default"`
}
