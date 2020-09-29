package http

import (
	pb "easymarket-go-server/app/api"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"easymarket-go-server/app/internal/server/middleware"
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
	// token校验
	engine = bm.DefaultServer(&cfg)

	engine.UseFunc(middleware.TokenValidate)
	
	pb.RegisterAppBMServer(engine, s)
	err = engine.Start()
	return
}
