// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"goTraining/Week04/internal/dao"
	"goTraining/Week04/internal/service"
	"goTraining/Week04/internal/server/grpc"
	"goTraining/Week04/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
