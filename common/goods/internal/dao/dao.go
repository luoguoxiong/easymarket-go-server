package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// Provider ...
var Provider = wire.NewSet(New, NewDB)

// Dao dao.
type Dao struct {
	db *gorm.DB
}

// New new a dao and return.
func New(db *gorm.DB) (d *Dao, cf func(), err error) {
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &Dao{
		db: db,
	}
	return
}
