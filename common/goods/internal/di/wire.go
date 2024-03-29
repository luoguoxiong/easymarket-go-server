// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"easymarket-go-server/common/goods/internal/dao"
	"easymarket-go-server/common/goods/internal/server/grpc"
	"easymarket-go-server/common/goods/internal/server/http"
	"easymarket-go-server/common/goods/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
