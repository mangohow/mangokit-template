package server

import (
	"github.com/mangohow/mangokit-template/api/helloworld"
	"github.com/mangohow/mangokit-template/internal/conf"
	"github.com/mangohow/mangokit-template/internal/service"
	"github.com/mangohow/mangokit/transport/httpwrapper"
	"github.com/sirupsen/logrus"
)

func NewHttpServer(serverConf *conf.Server,
	logger *logrus.Logger,
	greeterService *service.GreeterService) *httpwrapper.Server {

	server := httpwrapper.New(
		httpwrapper.WithAddr(serverConf.Http.Addr),
		httpwrapper.WithLogger(logger),
	)

	helloworld.RegisterGreeterHTTPService(server, greeterService)

	return server
}
