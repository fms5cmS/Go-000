// +build wireinject

package main

import (
	"database/sql"
	
	"goTraining/Week04/internal/biz"
	"goTraining/Week04/internal/dao"
	"goTraining/Week04/internal/service"
	"github.com/google/wire"
)

// wire.go

func InitHandler(db *sql.DB) service.Handler {
	wire.Build(service.NewHandler, biz.NewStringService, dao.NewRepo)
	return &service.HandlerImp{}
}
