// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"easymarketgoserve/app/internal/dao"
	"easymarketgoserve/app/internal/server/grpc"
	"easymarketgoserve/app/internal/server/http"
	"easymarketgoserve/app/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, grpc.New, http.New, NewApp))
}
