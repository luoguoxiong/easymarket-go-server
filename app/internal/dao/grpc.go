package dao

import (
	goods "easymarketgoserve/common/goods/api"

	"context"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	xtime "github.com/bilibili/kratos/pkg/time"
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
func NewGrpcClient() goods.DemoClient {
	var (
		service struct {
			Servers map[string]*ServerConfig
		}
	)
	err := paladin.Get("dao.toml").UnmarshalTOML(&service)
	if err != nil {
		panic(err)
	}
	config := service.Servers["goods-service"]
	cfg := &GRPCConf{
		Addr: config.Addr,
		Server: &warden.ClientConfig{
			Timeout: config.Timeout,
		},
		MaxReceiveMessageSize: config.MaxReceiveMessageSize,
	}

	cc, err := warden.NewClient(cfg.Server).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return goods.NewDemoClient(cc)
}
