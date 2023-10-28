package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Email     string     `                                json:"email"`
	Password  string     `gorm:"type:char(60)"            json:"password"`
	FirstName string     `                                json:"f_name"`
	LastName  string     `                                json:"l_name"`
	Tel       string     `                                json:"tel"`
	Address   []Address  `gorm:"foreignKey:UserID"`
	Carts     []Cart     `gorm:"foreignKey:UserID"`
	Trackings []Tracking `gorm:"foreignKey:UserID"`
	Reviews   []Review   `gorm:"foreignKey:UserID"`
}

type CreateUserRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
	Tel       string `json:"tel"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ResetPasswordRequest struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type UpdateInfoRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
	Tel       string `json:"tel"`
}
