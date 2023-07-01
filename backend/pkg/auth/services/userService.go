package service

import "github.com/mjyocca/go-auth/backend/pkg/auth/db"

type UserService struct {
	Db db.Handler
}

func (u *UserService) GetUser() {

}
