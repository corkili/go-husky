package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
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

func (dao *UserDao) FindByPhone(phone string) (user *entity.User, exist bool, err error)  {
	user = &entity.User{}
	err = dao.db.Where(&entity.User{Phone: phone}).First(user).Error
	if err != nil  {
		if gorm.IsRecordNotFoundError(err) {
			return nil, false, nil
		} else {
			logger.Error(err)
			return nil, false, errors.New("system error: cannot query user from db")
		}
	}
	return user, true, nil
}