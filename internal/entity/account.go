package entity

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type AccountType string

func (accountType AccountType) Name() string {
	return string (accountType)
}

func GetAccountTypeByName(name string) (AccountType, error) {
	if name == INCOME.Name() {
		return INCOME, nil
	} else if name == EXPENSE.Name() {
		return EXPENSE, nil
	} else {
		return "", errors.New("invalid name of AccountType")
	}
}

const (
	INCOME AccountType = "income"
	EXPENSE AccountType = "expense"
)

type Account struct {
	gorm.Model
	AccountTime *time.Time `gorm:"not null"`
	AccountType AccountType `gorm:"type:varchar(20);not null"`
	Money float64 `gorm:"not null"`
	Description string `gorm:"varchar(4096);not null"`
	AccountBooks []*AccountBook `gorm:"many2many:account_book_relation;"`
	User *User
	UserId uint
}

func (account *Account) GetId() uint {
	return account.ID
}
