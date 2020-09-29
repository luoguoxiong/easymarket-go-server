package dao

// import (
// 	"github.com/go-kratos/kratos/pkg/conf/paladin"
// 	_ "github.com/go-sql-driver/mysql" // init
// 	"github.com/jinzhu/gorm"
// )

// type cfg struct {
// 	Dsn          string
// 	Debug        bool
// 	MaxOpenConns int
// 	MaxIdleConns int
// }

// // NewDB ...
// func NewDB() (db *gorm.DB, cf func(), err error) {

// 	var (
// 		conf cfg
// 		ct   paladin.TOML
// 	)
// 	if err = paladin.Get("mysql.toml").Unmarshal(&ct); err != nil {
// 		return
// 	}
// 	if err = ct.Get("GoodsDB").UnmarshalTOML(&conf); err != nil {
// 		return
// 	}
// 	db, err = gorm.Open("mysql", conf.Dsn)
// 	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
// 	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
// 	db.LogMode(conf.Debug)
// 	cf = func() { db.Close() }
// 	return
// }
