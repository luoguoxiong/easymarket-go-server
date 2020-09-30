package libary

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	_ "github.com/go-sql-driver/mysql" //init
	"github.com/jinzhu/gorm"
)

type mysqlDb struct {
	Dsn          string
	Debug        bool
	MaxOpenConns int
	MaxIdleConns int
}

type mysqlConf struct {
	MysqlDb *mysqlDb
}

// MysqlInit mysql初始化
func MysqlInit(mysqlTomlName string) (db *gorm.DB) {
	mysqlConf := &mysqlConf{}
	if err := paladin.Get(mysqlTomlName).UnmarshalTOML(mysqlConf); err != nil {
		panic(err)
	}
	db, err := gorm.Open("mysql", mysqlConf.MysqlDb.Dsn)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(mysqlConf.MysqlDb.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysqlConf.MysqlDb.MaxOpenConns)
	db.LogMode(mysqlConf.MysqlDb.Debug)
	return
}
