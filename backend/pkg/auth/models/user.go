package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         int64     `gorm:"primaryKey" json:"-"`
	ExternalID uuid.UUID `gorm:"type:uuid;index" json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"` // to-do: add encryption
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ExternalID = uuid.New()
	return
}
