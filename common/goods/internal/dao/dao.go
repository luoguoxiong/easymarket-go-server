package dao

import (
	"easymarket-go-server/libary"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// Provider ...
var Provider = wire.NewSet(New)

// Dao dao.
type Dao struct {
	db *gorm.DB
}

// New new a dao and return.
func New() (d *Dao, err error) {
	db := libary.MysqlInit("mysql.toml")
	d = &Dao{
		db: db,
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
}
