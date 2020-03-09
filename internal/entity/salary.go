package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Salary struct {
	gorm.Model
	Predictive bool `gorm:"not null"`
	TimeOfSalary *time.Time `gorm:"not null"`
	TaxYear uint32 `gorm:"not null"`
	SalaryValue float64 `gorm:"not null"`
	Subsidy float64 `gorm:"not null"`
	Bonus float64 `gorm:"not null"`
	Remark string `gorm:"type:varchar(4096);not null"`
	User *User
	UserId uint
}