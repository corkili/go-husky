package dao

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go-husky/internal/db"
	"go-husky/internal/entity"
)

type CommonDao struct {
	db *gorm.DB
}

func (dao *CommonDao) init() {
	database, err := db.GetInstance().GetDB()
	if err != nil {
		logger.Error(err.Error())
	}
	dao.db = database
}

func (dao *CommonDao) CreateEntity(entity entity.Entity) (err error) {
	return dao.db.Create(entity).Error
}

func (dao *CommonDao) CreateEntities(entities ...entity.Entity) (successCount uint)  {
	successCount = 0
	for _, e := range entities {
		if err := dao.db.Create(e).Error; err != nil {
			logger.Warn(err.Error())
		} else {
			successCount++
		}
	}
	return successCount
}

func (dao *CommonDao) UpdateEntity(entity entity.Entity) (err error) {
	return dao.db.Save(entity).Error
}

func (dao *CommonDao) UpdateEntities(entities ...entity.Entity) (successCount uint) {
	successCount = 0
	for _, e := range entities {
		if err := dao.db.Save(e).Error; err != nil {
			logger.Warn(err.Error())
		} else {
			successCount++
		}
	}
	return successCount
}

func (dao *CommonDao) DeleteEntity(entity entity.Entity) (err error)  {
	if entity.GetId() > 0 {
		return dao.db.Delete(entity).Error
	} else {
		return errors.New("entity id not exists")
	}
}

func (dao *CommonDao) DeleteEntities(entities ...entity.Entity) (successCount uint)  {
	successCount = 0
	for _, e := range entities {
		if err := dao.DeleteEntity(e); err != nil {
			logger.Warn(err.Error())
		} else {
			successCount++
		}
	}
	return successCount
}


