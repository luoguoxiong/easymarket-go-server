package dao

import (
	"fmt"
	"time"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

var Provider = wire.NewSet(New, NewDB)

// Dao dao.
type Dao struct {
	db         *gorm.DB
	cache      *fanout.Fanout
	demoExpire int32
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
		db:         db,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.cache.Close()
}

// GetGoods ...
func (d *Dao) GetGoods() {
	var count int32
	d.db.Table("nideshop_goods").Count(&count)
	fmt.Println(count)
}
