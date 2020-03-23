package dao

import (
	"context"
	goods "easymarketgoserve/common/goods/api"
	"fmt"
	"time"

	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/database/sql"
	"github.com/bilibili/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/bilibili/kratos/pkg/time"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, NewRedis, NewDB)

// Dao dao interface
type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
	Ping2(ctx context.Context) (err error)
}

// dao dao.
type dao struct {
	db         *sql.DB
	redis      *redis.Redis
	cache      *fanout.Fanout
	demoExpire int32
	goodsGrpc  goods.DemoClient
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB) (d Dao, cf func(), err error) {
	return newDao(r, db)
}

func newDao(r *redis.Redis, db *sql.DB) (d *dao, cf func(), err error) {
	c := NewGrpcClient()
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &dao{
		db:         db,
		redis:      r,
		cache:      fanout.New("cache"),
		demoExpire: int32(time.Duration(cfg.DemoExpire) / time.Second),
		goodsGrpc:  c,
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *dao) Close() {
	d.cache.Close()
}

// Ping ping the resource.
func (d *dao) Ping(ctx context.Context) (err error) {
	return nil
}

func (d *dao) Ping2(ctx context.Context) (err error) {
	req := goods.HelloReq{
		Name: "2323",
	}
	res, _ := d.goodsGrpc.SayHelloURL(ctx, &req)
	fmt.Println(res)
	return nil
}
