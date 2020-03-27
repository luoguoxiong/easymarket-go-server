package http

import (
	pb "easymarketgoserve/app/api"
	"easymarketgoserve/app/internal/model"

	"github.com/bilibili/kratos/pkg/conf/paladin"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

var svc pb.AppServer

// New new a bm server.
func New(s pb.AppServer) (engine *bm.Engine, err error) {
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
	initRouter(engine)
	initBmRouter(engine, s)
	err = engine.Start()
	return
}

func initBmRouter(e *bm.Engine, s pb.AppServer) {
	// e.UseFunc(func(c *bm.Context) {
	// 	fmt.Println("232")
	// })
	pb.RegisterAppBMServer(e, s)
}

func initRouter(e *bm.Engine) {
	g := e.Group("/app")
	{
		g.GET("/start", howToStart)
	}
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
