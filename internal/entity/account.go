package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AccountType string

const (
	INCOME AccountType = "income"
	EXPENSE AccountType = "expense"
)

type Account struct {
	gorm.Model
	AccountBook AccountBook
	AccountBookId uint
	AlreadyGenerated bool `gorm:"not null"`
	AccountTime *time.Time `gorm:"not null"`
	AccountType AccountType `gorm:"type:varchar(20);not null"`
	Money float64 `gorm:"not null"`
	Description string `gorm:"varchar(4096);not null"`
}
