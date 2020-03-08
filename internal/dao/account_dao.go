package dao

import (
	"fmt"
	"go-husky/internal/api"
	"go-husky/internal/entity"
	"strings"
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
	dao.db.Preload("AccountBooks").First(account, id)
	return account
}

func (dao *AccountDao) FindAllByUser(user *entity.User) []*entity.Account {
	var accounts = make([]*entity.Account, 0)
	dao.db.Where(&entity.Account{
		UserId: user.ID,
	}).Preload("AccountBooks").Find(&accounts)
	return accounts
}

func (dao *AccountDao) CountByUser(user *entity.User) (count uint) {
	dao.db.Model(&entity.Account{}).Where(&entity.Account{
		UserId: user.ID,
	}).Count(&count)
	return count
}

func (dao *AccountDao) FindByPaginationAndSort(pageSize uint, currentPage uint, sort *api.Sort, filters []api.Filter, user *entity.User) []*entity.Account {
	db := dao.db.Where(&entity.Account{
		UserId: user.ID,
	})
	for _, filter := range filters {
		switch strings.ToLower(filter.Operation) {
		case "and":
			db = db.Where(map[string]interface{}{filter.Field: filter.Value})
		case "or":
			db = db.Or(map[string]interface{}{filter.Field: filter.Value})
		case "not":
			db = db.Not(map[string]interface{}{filter.Field: filter.Value})
		case "like":
			db = db.Where(fmt.Sprintf("%s like ?", filter.Field), fmt.Sprintf("%%%s%%", filter.Value.(string)))
		}
	}
	if sort != nil && len(sort.SortRules) > 0 {
		for _, rule := range sort.SortRules {
			db = db.Order(fmt.Sprintf("%s %s", rule.Field, rule.Method))
		}
	}
	var accounts = make([]*entity.Account, 0)
	db.Offset(pageSize * (currentPage - 1)).Limit(pageSize).Preload("AccountBooks").Find(&accounts)
	return accounts
}


