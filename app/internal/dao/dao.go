package dao

import (
	goods "easymarket-go-server/common/goods/api"
	topic "easymarket-go-server/common/topic/api"

	"github.com/go-kratos/kratos/pkg/cache/redis"
	"github.com/go-kratos/kratos/pkg/database/sql"
	"github.com/go-kratos/kratos/pkg/sync/pipeline/fanout"

	"github.com/google/wire"
)

// Provider ...
var Provider = wire.NewSet(New, NewRedis, NewDB)

// Dao dao.
type Dao struct {
	db          *sql.DB
	redis       *redis.Redis
	cache       *fanout.Fanout
	demoExpire  int32
	goodsClient goods.GoodsClient
	topicClient topic.TopicClient
}

// New new a dao and return.
func New(r *redis.Redis, db *sql.DB) (d *Dao, cf func(), err error) {
	// g := NewGrpcConfig()
	d = &Dao{
		db:          db,
		redis:       r,
		cache:       fanout.New("cache"),
		// goodsClient: g.NewGoodsClient(),
		// topicClient: g.NewTopicClient(),
	}
	cf = d.Close
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.cache.Close()
}
