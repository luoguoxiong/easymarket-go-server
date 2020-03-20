// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"easymarketgoserve/common/goods/internal/dao"
	"easymarketgoserve/common/goods/internal/service"
	"easymarketgoserve/common/goods/internal/server/grpc"
	"easymarketgoserve/common/goods/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
