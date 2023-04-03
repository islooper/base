package sharedconfig

import (
	"fmt"
	"gitee.com/hzxkd/base/dao"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DBConf struct {
	Host        string `json:",default=0.0.0.0"`
	Port        int    `json:",default=3306"`
	Name        string
	User        string `json:",default=root"`
	Passwd      string
	Prefix      string `json:",default=t_"`
	MaxIdleNum  int    `json:",default=10"`  // 最大空闲数
	MaxOpenNum  int    `json:",default=100"` // 最大连接数
	MaxLifeTime int    `json:",default=30"`  // 最长生存周期
}

// NewDB 初始化db连接
func NewDB(conf DBConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User,
		conf.Passwd, conf.Host, conf.Port, conf.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return db, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return db, err
	}

	sqlDB.SetMaxIdleConns(conf.MaxIdleNum)
	sqlDB.SetMaxOpenConns(conf.MaxOpenNum)
	sqlDB.SetConnMaxLifetime(time.Duration(conf.MaxLifeTime) * time.Minute)
	return db, err
}


// NewDBCollection 通过一组db conf初始化db资源
func NewDBCollection(confGroup map[string]DBConf) (dao.DBCollector, error) {
	collection := &dao.DBCollection{
		Collection: map[string]*gorm.DB{},
	}

	for label, conf := range confGroup {
		if db, err := NewDB(conf); err != nil {
			return collection, err
		} else {
			collection.RegisterDB(label, db)
		}
	}

	return collection, nil
}
