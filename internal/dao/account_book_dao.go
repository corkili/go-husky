package dao

import (
	"go-husky/internal/entity"
	"sync"
)

type AccountBookDao struct {
	CommonDao
}

var accountBookDao *AccountBookDao
var accountBookDaoOnce sync.Once

func GetAccountBookDao() (dao *AccountBookDao) {
	accountBookDaoOnce.Do(func() {
		accountBookDao = &AccountBookDao{}
	})
	accountBookDao.init()
	return accountBookDao
}

func (dao *AccountBookDao) FindById(id uint) (accountBook *entity.AccountBook) {
	accountBook = &entity.AccountBook{}
	dao.db.First(accountBook, id)
	return accountBook
}

