// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"easymarket-go-server/app/internal/dao"
	"easymarket-go-server/app/internal/server/grpc"
	"easymarket-go-server/app/internal/server/http"
	"easymarket-go-server/app/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
