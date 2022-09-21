// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"lusun/internal/biz"
	"lusun/internal/conf"
	"lusun/internal/data"
	"lusun/internal/server"
	"lusun/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	mysqlRepo := data.NewMysqlRepo()
	articleRepo := data.NewArticle(mysqlRepo)
	articleUseCase := biz.NewArticleUseCase(articleRepo)
	greeterService := service.NewGreeterService(greeterUsecase, articleUseCase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
