package dao

import (
	"go-husky/internal/entity"
	"sync"
)

type SalaryDao struct {
	CommonDao
}

var salaryDao *SalaryDao
var salaryDaoOnce sync.Once

func GetSalaryDao() (dao *SalaryDao) {
	salaryDaoOnce.Do(func() {
		salaryDao = &SalaryDao{}
	})
	salaryDao.init()
	return salaryDao
}

func (dao *SalaryDao) FindById(id uint) (salary *entity.Salary) {
	salary = &entity.Salary{}
	dao.db.First(salary, id)
	return salary
}

