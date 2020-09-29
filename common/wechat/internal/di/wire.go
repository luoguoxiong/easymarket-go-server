// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"easymarket-go-server/common/wechat/internal/dao"
	"easymarket-go-server/common/wechat/internal/server/grpc"
	"easymarket-go-server/common/wechat/internal/server/http"
	"easymarket-go-server/common/wechat/internal/service"

	"github.com/google/wire"
)

// InitApp ...
//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
