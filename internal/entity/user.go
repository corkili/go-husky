package entity

import (
	"github.com/jinzhu/gorm"
)

type UserRole uint16

const (
	ADMIN UserRole = 0
	USER  UserRole = 1
)

type User struct {
	gorm.Model
	Phone    string 	`gorm:"type:varchar(20);not null;index:phone_index"`
	Password string 	`gorm:"type:varchar(512);not null"`
	Username string		`gorm:"type:varchar(100);not null"`
	Roll		 UserRole `gorm:"not null"`
}

func (user *User) GetId() uint {
	return user.ID
}
