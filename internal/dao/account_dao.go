package dao

import (
	"go-husky/internal/entity"
	"sync"
)

type AccountDao struct {
	CommonDao
}

var accountDao *AccountDao
var accountDaoOnce sync.Once

func GetAccountDao() (dao *AccountDao) {
	accountDaoOnce.Do(func() {
		accountDao = &AccountDao{}
	})
	accountDao.init()
	return accountDao
}

func (dao *AccountDao) FindById(id uint) (account *entity.Account) {
	account = &entity.Account{}
	dao.db.First(account, id)
	return account
}

