package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
  Email string `json:"email"`
  Password string `json:"password"`
  FirstName string `json:"f_name"`
  LastName string `json:"l_name"`
  Tel string `json:"tel"`
  Address []Address `gorm:"foreignKey:UserID"`
}

type LoginRequest struct {
 Email    string `json:"email"`
 Password string `json:"password"`
}

type LoginResponse struct {
 Token string `json:"token"`
}

