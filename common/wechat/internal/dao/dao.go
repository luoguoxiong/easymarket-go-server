package dao

import (
	"easymarket-go-server/libary"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	gRedis "github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" //init
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// Provider ...
var Provider = wire.NewSet(New)

// Dao dao.
type Dao struct {
	db     *gorm.DB
	wechat *wechat
	redis  *gRedis.Client
}

// wechat 微信配置
type wechat struct {
	Appid  string
	Secret string
}

type weChatConf struct {
	Wechat *wechat
}

// New new a dao and return.
func New() (d *Dao, cf func(), err error) {

	weChatConf := &weChatConf{}

	if err = paladin.Get("application.toml").UnmarshalTOML(&weChatConf); err != nil {
		panic(err)
	}

	redis := libary.RedisInit("redis.toml")
	db := libary.MysqlInit("mysql.toml")

	d = &Dao{
		db,
		weChatConf.Wechat,
		redis,
	}

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
	d.redis.Close()
}
