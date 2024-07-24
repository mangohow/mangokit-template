package server

import (
	"github.com/mangohow/mangokit-template/api/greeter"
	"github.com/mangohow/mangokit-template/internal/conf"
	"github.com/mangohow/mangokit-template/internal/middleware"
	"github.com/mangohow/mangokit-template/internal/service"
	"github.com/mangohow/mangokit/transport/http"
	"github.com/sirupsen/logrus"
)

func NewHttpServer(serverConf *conf.Server,
	logger *logrus.Logger,
	greeterService *service.GreeterService) *http.Server {

	server := http.New(
		http.WithAddr(serverConf.Http.Addr),
		http.WithLogger(logger),
	)

	// 添加middleware
	server.Middleware(
		middleware.RunTime(logger), // 统计运行时间
		middleware.Logger(logger),  // 打印日志
		middleware.Validator(),     // 校验参数
	)

	greeter.RegisterGreeterHTTPService(server, greeterService)

	return server
}
