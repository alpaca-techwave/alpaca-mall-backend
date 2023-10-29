package repository

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

func FindByUserCredentials(email string, password string) (*models.User, error) {
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

func FindByAdminCredentials(email string, password string) (*models.Admin, error) {
	var admin models.Admin
	result := config.Database.Where(&models.Admin{Email: email}).Find(&admin)
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err == nil && result.RowsAffected != 0 {
		return &models.Admin{
			ID:    admin.ID,
			Email: admin.Email,
		}, nil
	}
	return nil, errors.New("Unauthorize")
}
