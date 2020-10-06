package dao

import (
	goods "easymarket-go-server/common/goods/api"
	topic "easymarket-go-server/common/topic/api"
	wechat "easymarket-go-server/common/wechat/api"
	"easymarket-go-server/libary"
	gRedis "github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

// Provider ...
var Provider = wire.NewSet(New)

// Dao dao.
type Dao struct {
	db           *gorm.DB
	redis        *gRedis.Client
	goodsClient  goods.GoodsClient
	topicClient  topic.TopicClient
	wechatClient wechat.WeChatClient
}

// New new a dao and return.
func New() (d *Dao, err error) {
	db := libary.MysqlInit("mysql.toml")
	redis := libary.RedisInit("redis.toml")
	grpc := NewGrpcConfig()

	d = &Dao{
		db:           db,
		redis:        redis,
		wechatClient: grpc.NewWechatClient(),
		topicClient:  grpc.NewTopicClient(),
		goodsClient:  grpc.NewGoodsClient(),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.db.Close()
	d.redis.Close()
}
