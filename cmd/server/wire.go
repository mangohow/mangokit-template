//go:generate wire
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mangohow/mangokit-template/internal/conf"
	"github.com/mangohow/mangokit-template/internal/dao"
	"github.com/mangohow/mangokit-template/internal/server"
	"github.com/mangohow/mangokit-template/internal/service"
	"github.com/mangohow/mangokit/transport/httpwrapper"
	"github.com/sirupsen/logrus"
)

func newApp(serverConf *conf.Server, dataConf *conf.Data, logger *logrus.Logger) (*httpwrapper.Server, func(), error) {
	panic(wire.Build(dao.ProviderSet, service.ProviderSet, server.NewHttpServer))
}
