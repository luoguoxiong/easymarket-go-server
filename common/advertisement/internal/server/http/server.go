package http

import (
	pb "easymarket-go-server/common/goods/api"
	"easymarket-go-server/common/goods/internal/model"
	"fmt"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svc pb.GoodsServer

// New new a bm server.
func New(s pb.GoodsServer) (engine *bm.Engine, err error) {
	var (
		cfg bm.ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svc = s
	engine = bm.DefaultServer(&cfg)
	pb.RegisterGoodsBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

func initRouter(e *bm.Engine) {
	g := e.Group("/goods")
	{
		g.GET("/start", howToStart)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "",
		Lis:   make([]string, 0),
		Num:   0,
	}
	fmt.Println(k)
	c.JSON(k, nil)
}
