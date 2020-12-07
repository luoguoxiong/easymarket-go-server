package dao

import (
	"context"
	advertisement "easymarket-go-server/common/advertisement/api"
	goods "easymarket-go-server/common/goods/api"
	topic "easymarket-go-server/common/topic/api"
	wechat "easymarket-go-server/common/wechat/api"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	xtime "github.com/go-kratos/kratos/pkg/time"
	"google.golang.org/grpc"
)

// GrpcServiceConfig grpc服务配置
type GrpcServiceConfig struct {
	Servers map[string]*ServerConfig
}

// ServerConfig 配置文件
type ServerConfig struct {
	Addr                  string
	Timeout               xtime.Duration
	MaxReceiveMessageSize int
}

// GRPCConf grpcclient配置
type GRPCConf struct {
	Server                *warden.ClientConfig
	Addr                  string
	MaxReceiveMessageSize int
}

// NewGrpcConfig 初始化配置文件
func NewGrpcConfig() *GrpcServiceConfig {
	var service GrpcServiceConfig
	err := paladin.Get("dao.toml").UnmarshalTOML(&service)
	if err != nil {
		panic(err)
	}
	return &service
}

// setConfig 初始化grpc连接
func setConfig(g *GrpcServiceConfig, serverName string) (conn *grpc.ClientConn, err error) {
	config := g.Servers[serverName]
	cfg := &GRPCConf{
		Addr: config.Addr,
		Server: &warden.ClientConfig{
			Timeout: config.Timeout,
		},
		MaxReceiveMessageSize: config.MaxReceiveMessageSize,
	}
	conn, err = warden.NewClient(cfg.Server).Dial(context.Background(), cfg.Addr)
	if err != nil {
		panic(err)
	}
	return
}

// NewGoodsClient 实例化goods-grpc服务客户端
func (g *GrpcServiceConfig) NewGoodsClient() goods.GoodsClient {
	c, _ := setConfig(g, "goods-service")
	return goods.NewGoodsClient(c)
}

// NewTopicClient 实例化topic-grpc服务客户端
func (g *GrpcServiceConfig) NewTopicClient() topic.TopicClient {
	c, _ := setConfig(g, "topic-service")
	return topic.NewTopicClient(c)
}

// NewWechatClient 实例化wechat-grpc服务客户端
func (g *GrpcServiceConfig) NewWechatClient() wechat.WeChatClient {
	c, _ := setConfig(g, "wechat-service")
	return wechat.NewWeChatClient(c)
}

// NewAdsClient 实例化advertisement-grpc服务客户端
func (g *GrpcServiceConfig) NewAdsClient() advertisement.AdvertisementClient {
	c, _ := setConfig(g, "advertisement-service")
	return advertisement.NewAdvertisementClient(c)
}
