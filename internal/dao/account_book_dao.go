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

func (dao *AccountBookDao) FindAllByUser(user *entity.User) []*entity.AccountBook {
	var accountBooks = make([]*entity.AccountBook, 0)
	dao.db.Where(&entity.AccountBook{
		UserId: user.ID,
	}).Find(&accountBooks)
	return accountBooks
}

