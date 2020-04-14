// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"easymarket-go-server/common/topic/internal/dao"
	"easymarket-go-server/common/topic/internal/server/grpc"
	"easymarket-go-server/common/topic/internal/server/http"
	"easymarket-go-server/common/topic/internal/service"

	"github.com/google/wire"
)

// InitApp ...
//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
