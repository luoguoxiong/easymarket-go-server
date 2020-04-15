package dao

import (
	goods "easymarket-go-server/common/goods/api"
	"fmt"

	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	xtime "github.com/go-kratos/kratos/pkg/time"
)

// GRPCConf config for client
type GRPCConf struct {
	Server                *warden.ClientConfig
	Addr                  string
	MaxReceiveMessageSize int
}

// ServerConfig ...
type ServerConfig struct {
	Addr                  string
	Timeout               xtime.Duration
	MaxReceiveMessageSize int
}

// NewGrpcClient 实例化Grpc
func NewGrpcClient() goods.GoodsClient {
	var (
		service struct {
			Servers map[string]*ServerConfig
		}
	)
	err := paladin.Get("dao.toml").UnmarshalTOML(&service)
	if err != nil {
		panic(err)
	}
	fmt.Println(service.Servers)
	goodsConfig := service.Servers["goods-service"]
	cfg := &GRPCConf{
		Addr: goodsConfig.Addr,
		Server: &warden.ClientConfig{
			Timeout: goodsConfig.Timeout,
		},
		MaxReceiveMessageSize: goodsConfig.MaxReceiveMessageSize,
	}
	topicConfig := service.Servers["topic-service"]
	fmt.Println(topicConfig)
	cc, err := warden.NewClient(cfg.Server).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return goods.NewGoodsClient(cc)
}
