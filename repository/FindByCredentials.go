package repository

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

func FindByCredentials(email string, password string) (*models.User, error) {
	var user models.User
	result := config.Database.Where(&models.User{Email: email}).Find(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil && result.RowsAffected != 0 {
		return &models.User{
			ID:    user.ID,
			Email: user.Email,
		}, nil
	}
	return nil, errors.New("Unauthorize")
}
