package repository

import (
	"errors"

	"github.com/alpaca-techwave/alpaca-mall-backend/config"
	"github.com/alpaca-techwave/alpaca-mall-backend/models"
)

func FindByCredentials(email string, password string) (*models.User, error) {
  var user models.User
  result := config.Database.Where(&models.User{Email: email,Password: password }).Find(&user)
  if result.RowsAffected != 0 {
    return &models.User {
      ID: user.ID,
      Email: user.Email,
    },nil
  }
  return nil, errors.New("user not found")
}
