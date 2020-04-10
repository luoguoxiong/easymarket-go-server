package dao

import (
	"context"
	goods "easymarket-go-server/common/goods/api"
	"time"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"
	xtime "github.com/go-kratos/kratos/pkg/time"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, NewRedis, NewDB)

// Dao dao.
type Dao struct {
	db         *sql.DB
	redis      *redis.Redis
	cache      *fanout.Fanout
	demoExpire int32
	goodsGrpc  goods.GoodsClient
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB) (d *Dao, cf func(), err error) {
	c := NewGrpcClient()
	var cfg struct {
		DemoExpire xtime.Duration
	}
	if err = paladin.Get("application.toml").UnmarshalTOML(&cfg); err != nil {
		return
	}
	d = &Dao{
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
func (d *Dao) Close() {
	d.cache.Close()
}

// GetGoodsList Ping2
func (d *Dao) GetGoodsList(ctx context.Context, req *goods.GoodsReq) (res *goods.GoodsListRes, err error) {
	res, err = d.goodsGrpc.GetGoodsList(ctx, req)
	return
}

// GetGoodsDetail Ping2
func (d *Dao) GetGoodsDetail(ctx context.Context, req *goods.GoodsDetailReq) (res *goods.GoodsRes, err error) {
	res, err = d.goodsGrpc.GetGoodsDetail(ctx, req)
	return
}
