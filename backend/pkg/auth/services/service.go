package service

import "github.com/mjyocca/go-auth/backend/pkg/auth/db"

type service struct {
	handler db.Handler
	UserService
}

func NewService(h db.Handler) *service {
	return &service{
		handler:     h,
		UserService: &userService{h},
	}
}
