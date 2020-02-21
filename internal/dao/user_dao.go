package dao

import (
	"go-husky/internal/entity"
	"sync"
)

type UserDao struct {
	CommonDao
}

var userDao *UserDao
var userDaoOnce sync.Once

func GetUserDao() (dao *UserDao) {
	userDaoOnce.Do(func() {
		userDao = &UserDao{}
	})
	userDao.init()
	return userDao
}

func (dao *UserDao) FindById(id uint) (user *entity.User) {
	user = &entity.User{}
	dao.db.First(user, id)
	return user
}

