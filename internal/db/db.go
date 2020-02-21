package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-husky/internal/entity"
	"go-husky/internal/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type DataSourceConf struct {
	User string `yaml:"user"`
	Password string `yaml:"password"`
	DBName string `yaml:"dbname"`
	Host string `yaml:"host"`
	Options string `yaml:"options"`
}

type DatabaseConnection struct {
	db *gorm.DB
}

const dataSourceConfFile = "./conf/db.yaml"
var logger = log.GetLogger()

var dsConf *DataSourceConf

var instance *DatabaseConnection
var once sync.Once

var entities = []interface{} {
	&entity.User{},
	&entity.AccountBook{},
	&entity.Account{},
	&entity.Salary{},
}

func GetInstance() *DatabaseConnection {
	once.Do(func() {
		instance = &DatabaseConnection{}
	})
	return instance
}

func (dc *DatabaseConnection) Init() (err error)  {
	dc.db, err = getDB()
	if err != nil {
		logger.Error(err.Error())
	} else {
		dc.db.AutoMigrate(entities...)
	}
	return err
}

func (dc *DatabaseConnection) GetDB() (dbConn *gorm.DB, err error)  {
	if dc.db == nil {
		err = dc.Init()
		if err != nil {
			logger.Error(err.Error())
		}
	}
	return dc.db, err
}

func (dc *DatabaseConnection) Close() {
	if dc.db != nil {
		err := dc.db.Close()
		if err != nil {
			logger.Warn(err.Error())
		}
		dc.db = nil
	}
}

func getConf() (conf *DataSourceConf, error error) {
	if dsConf == nil {
		confFile, err := ioutil.ReadFile(dataSourceConfFile)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		dsConf = new(DataSourceConf)
		err = yaml.Unmarshal(confFile, dsConf)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
	}
	return dsConf, nil
}

func checkConf(conf *DataSourceConf) (err error) {
	if conf == nil {
		return errors.New("data source conf is nil")
	}
	if conf.User == "" {
		return errors.New("\"user\", in data source, is empty")
	}
	if conf.Password == "" {
		return errors.New("\"password\", in data source, is empty")
	}
	if conf.DBName == "" {
		return errors.New("\"dbname\", in data source, is empty")
	}
	return nil
}

func getDB() (db *gorm.DB, error error) {
	conf, err := getConf()
	if err != nil {
		 logger.Error(err.Error())
		 return nil, err
	}
	err = checkConf(conf)
	if err != nil {
		 logger.Error(err.Error())
	}
	var args = fmt.Sprintf("%s:%s@(%s)/%s?%s",
		conf.User, conf.Password, conf.Host, conf.DBName, conf.Options)
	fmt.Printf("try to open mysql db: %s\n", args)
	database, err := gorm.Open("mysql", args)
	if err != nil {
		logger.Error(err.Error())
	}
	return database, err
}
