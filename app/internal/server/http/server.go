package http

import (
	pb "easymarket-go-server/app/api"
	"easymarket-go-server/app/internal/server/middleware"
	"easymarket-go-server/app/internal/service"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/ecode"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

// New new a bm server.
func New(s *service.Service) (engine *bm.Engine, err error) {
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
	engine = bm.DefaultServer(&cfg)
	// token校验
	engine.UseFunc(middleware.TokenValidate(s))

	pb.RegisterAppBMServer(engine, s)
	ecode.Register(errcode)
	err = engine.Start()
	return
}
