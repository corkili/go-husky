package entity

import "github.com/jinzhu/gorm"

type AccountBook struct {
	gorm.Model
	Name string `gorm:"type:varchar(128);not null"`
	Description string `gorm:"type:varchar(4096);not null"`
	User User
	UserId uint
}
