package service

import (
	"github.com/jinzhu/gorm"
	"go-husky/internal/api"
	"go-husky/internal/dao"
	"go-husky/internal/entity"
	"sync"
	"time"
)

const datetimeFormat = "2006-01-02 15:04:05"

type AccountBookService struct {
	accountBookDao *dao.AccountBookDao
	accountDao *dao.AccountDao
	onlineAccountBooks map[uint]*entity.AccountBook
}

var accountBookService *AccountBookService
var accountBookServiceOnce sync.Once

func GetAccountBookService() (controller *AccountBookService) {
	accountBookServiceOnce.Do(func() {
		accountBookService = &AccountBookService{}
		accountBookService.onlineAccountBooks = make(map[uint]*entity.AccountBook)
		accountBookService.accountBookDao = dao.GetAccountBookDao()
		accountBookService.accountDao = dao.GetAccountDao()
	})
	return accountBookService
}

func (s *AccountBookService) retrieveAccountBooks(user *entity.User) []*entity.AccountBook {
	return s.accountBookDao.FindAllByUser(user)
}

func (s *AccountBookService) getAccountBookByName(accountBooks []*entity.AccountBook, name string) *entity.AccountBook {
	for _, accountBook := range accountBooks {
		if accountBook.Name == name {
			return accountBook
		}
	}
	return nil
}

func (s *AccountBookService) buildAccountRspData(account *entity.Account) map[string]interface{} {
	accountBookArr := make([]map[string]interface{}, 0)
	for _, accountBook := range account.AccountBooks {
		accountBookArr = append(accountBookArr, s.buildSimpleAccountBookRspData(accountBook))
	}
	return map[string]interface{}{
		"id": account.ID,
		"datetime": account.AccountTime.Format(datetimeFormat),
		"type": account.AccountType.Name(),
		"money": account.Money,
		"description": account.Description,
		"accountBooks": accountBookArr,
	}
}

func (s *AccountBookService) buildSimpleAccountBookRspData(accountBook *entity.AccountBook) map[string]interface{} {
	return map[string]interface{}{
		"id": accountBook.ID,
		"name": accountBook.Name,
		"description": accountBook.Description,
	}
}

func (s *AccountBookService) SaveAccount(req *api.CreateOrUpdateAccountReq, user *entity.User) *api.Response  {

	var account = &entity.Account{
		Model: gorm.Model{
			ID: req.Id,
		},
		AccountTime:  nil,
		AccountType:  "",
		Money:        0,
		Description:  "",
		AccountBooks: nil,
		User:         nil,
		UserId:       0,
	}

	if _, rsp := userService.CheckIfUserLogin(user); rsp != nil {
		return rsp
	}

	account.User, account.UserId = user, user.ID

	if req.Datetime == "" {
		return BuildResponse(api.RspCodeAccountBookErrorAccountTimeEmpty,
			"datetime is empty", "账目时间不能为空", nil)
	}

	accountTime, err := time.ParseInLocation(datetimeFormat, req.Datetime, time.Local)
	if err != nil {
		logger.Error(err.Error())
		return BuildResponse(api.RspCodeAccountBookErrorAccountTimeInvalid,
			"datetime is invalid", "账目时间非法", nil)
	}
	account.AccountTime = &accountTime

	account.AccountType, err = entity.GetAccountTypeByName(req.Type)
	if err != nil {
		logger.Error(err.Error())
		return BuildResponse(api.RspCodeAccountBookErrorAccountTypeInvalid,
			"account type is invalid", "账目类型非法", nil)
	}

	if req.Money < 0 {
		return BuildResponse(api.RspCodeAccountBookErrorMoneyNegative,
			"money is negative", "账目金额必须大于或等于0", nil)
	}
	account.Money = req.Money

	if req.Description == "" {
		return BuildResponse(api.RspCodeAccountBookErrorDescriptionEmpty,
			"description is empty", "账目描述不能为空", nil)
	}
	account.Description = req.Description

	if req.AccountBooks == nil || len(req.AccountBooks) == 0  {
		return BuildResponse(api.RspCodeAccountBookErrorAccountBooksEmpty,
			"accountBooks is empty", "所属账簿不能为空", nil)
	}
	accountBooks := s.retrieveAccountBooks(user)
	for _, accountBook := range req.AccountBooks {
		accountBookEntity := s.getAccountBookByName(accountBooks, accountBook.Name)
		if accountBookEntity != nil {
			account.AccountBooks = append(account.AccountBooks, accountBookEntity)
		} else {
			accountBookEntity = &entity.AccountBook{
				Name:        accountBook.Name,
				Description: accountBook.Description,
				User:        user,
				UserId:      user.ID,
			}
			accountBooks = append(accountBooks, accountBookEntity)
			account.AccountBooks = append(account.AccountBooks, accountBookEntity)
		}
	}
	if account.ID == 0 {
		err = accountBookService.accountDao.CreateEntity(account)
	} else {
		err = accountBookService.accountDao.UpdateEntity(account)
	}
	if err != nil {
		logger.Error(err.Error())
		return BuildResponse(api.RspCodeDbError, "insert or update db error", "数据更新异常", nil)
	}
	return BuildResponse(api.RspCodeSuccess, "success", "操作成功", s.buildAccountRspData(account))
}

func (s *AccountBookService) RetrieveAccount(req *api.RetrieveAccountReq, user *entity.User) *api.Response {
	if _, rsp := userService.CheckIfUserLogin(user); rsp != nil {
		return rsp
	}
	accountArr := make([]map[string]interface{}, 0)
	accounts := s.accountDao.FindAllByUser(user)
	for _, account := range accounts {
		accountArr = append(accountArr, s.buildAccountRspData(account))
	}
	return BuildResponse(api.RspCodeSuccess, "success", "获取成功", map[string]interface{}{
		"accounts": accountArr,
	})
}

func (s *AccountBookService) RetrieveAccountBook(req *api.RetrieveAccountBookReq, user *entity.User) *api.Response {
	if _, rsp := userService.CheckIfUserLogin(user); rsp != nil {
		return rsp
	}
	accountBookArr := make([]map[string]interface{}, 0)
	accountBooks := s.accountBookDao.FindAllByUser(user)
	for _, accountBook := range accountBooks {
		accountBookArr = append(accountBookArr, s.buildSimpleAccountBookRspData(accountBook))
	}
	return BuildResponse(api.RspCodeSuccess, "success", "获取成功", map[string]interface{}{
		"accountBooks": accountBookArr,
	})
}

func (s *AccountBookService) DeleteAccount(req *api.DeleteAccountReq, user *entity.User) *api.Response {
	if _, rsp := userService.CheckIfUserLogin(user); rsp != nil {
		return rsp
	}
	accounts := make([]entity.Entity, 0)
	for _, id := range req.Ids {
		accounts = append(accounts, &entity.Account{
			Model: gorm.Model{
				ID: id,
			},
		})
	}
	successIds := s.accountDao.DeleteEntities(accounts...)
	return BuildResponse(api.RspCodeSuccess, "success", "删除成功", map[string]interface{}{
		"successIds": successIds,
	})
}

