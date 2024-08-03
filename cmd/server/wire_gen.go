// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/mangohow/mangokit-template/internal/conf"
	"github.com/mangohow/mangokit-template/internal/dao"
	"github.com/mangohow/mangokit-template/internal/server"
	"github.com/mangohow/mangokit-template/internal/service"
	"github.com/mangohow/mangokit/transport/http"
	"github.com/sirupsen/logrus"
)

// Injectors from wire.go:

func newApp(serverConf *conf.Server, dataConf *conf.Data, logger *logrus.Logger) (*http.Server, func(), error) {
	db, cleanup, err := dao.NewFakeMysqlClient(dataConf)
	if err != nil {
		return nil, nil, err
	}
	greeterDao := dao.NewGreeterDao(db)
	greeterService := service.NewGreeterService(greeterDao, logger)
	httpServer := server.NewHttpServer(serverConf, logger, greeterService)
	return httpServer, func() {
		cleanup()
	}, nil
}
