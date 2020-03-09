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
	if name == Income.Name() {
		return Income, nil
	} else if name == Expense.Name() {
		return Expense, nil
	} else if name == Loan.Name() {
		return Loan, nil
	} else if name == RepayLoan.Name() {
		return RepayLoan, nil
	} else {
		return "", errors.New("invalid name of AccountType")
	}
}

const (
	Income    AccountType = "income"
	Expense   AccountType = "expense"
	Loan      AccountType = "loan"
	RepayLoan AccountType = "repay_loan"
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
