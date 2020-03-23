package dao

import (
	_ "github.com/go-sql-driver/mysql" // init
	"github.com/jinzhu/gorm"
)

// NewDB ...
func NewDB() (db *gorm.DB, cf func(), err error) {
	db, err = gorm.Open("mysql", "root:ectripaes2013@tcp(202.96.155.121)/nideshop?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	cf = func() { db.Close() }
	return
}
