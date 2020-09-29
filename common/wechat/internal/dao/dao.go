package dao

import (
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql" //init
)

// Provider ...
var Provider = wire.NewSet(New)


// mysqlDb mysql配置
type mysqlDb struct {
	Dsn          string
	Debug        bool
	MaxOpenConns int
	MaxIdleConns int
}

// wechat 微信配置
type wechat struct {
	Appid  string
	Secret string
}

// Dao dao.
type Dao struct {
	db     *gorm.DB
	wechat *wechat
}

// New new a dao and return.
func New() (d *Dao, cf func(), err error) {
	
	var (
		weChatConf struct {
			Wechat *wechat
		}
		mysqlConf struct{
			MysqlDb *mysqlDb
		}
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&weChatConf); err != nil {
		panic(err)
	}

	if err = paladin.Get("mysql.toml").UnmarshalTOML(&mysqlConf); err != nil {
		panic(err)
	}

	db, err := gorm.Open("mysql", mysqlConf.MysqlDb.Dsn)
	db.DB().SetMaxIdleConns(mysqlConf.MysqlDb.MaxIdleConns)
	db.DB().SetMaxOpenConns(mysqlConf.MysqlDb.MaxOpenConns)
	db.LogMode(mysqlConf.MysqlDb.Debug)

	d = &Dao{
		db:     db,
		wechat: weChatConf.Wechat,
	}

	return
}

// Close close the resource.
func (d *Dao) Close() {
	fmt.Println(111)
	d.db.Close()
}
