package service

import (
	"github.com/mjyocca/go-auth/backend/pkg/auth/db"
	"github.com/mjyocca/go-auth/backend/pkg/auth/models"
)

type UserService interface {
	GetUser()
}

type userService struct {
	handler db.Handler
}

func (u *userService) GetUser() {
	var user models.User
	if result := u.handler.DB.Where(&models.User{Email: ""}).First(&user); result.Error == nil {

	}
}
