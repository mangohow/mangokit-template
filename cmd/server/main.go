package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mangohow/mangokit-template/internal/conf"
	"github.com/mangohow/mangokit-template/pkg/logger"
	"github.com/mangohow/mangokit/proc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// 加载配置文件
	cfg, err := loadConfig()
	if err != nil {
		panic(fmt.Sprintf("load config err=%v", err))
	}

	// 创建logger
	err = logger.InitLogger(cfg.Logger)
	if err != nil {
		panic(fmt.Sprintf("create logger err=%v", err))
	}

	// 创建server
	server, cleanup, err := newApp(cfg.Server, cfg.Data, logger.Logger())
	if err != nil {
		panic(fmt.Sprintf("init err=%v", err))
	}

	// 设置信号处理器
	ctx := proc.SetupSignalHandler()

	// 启动server
	go func() {
		if err := server.Start(); err != nil {
			panic(fmt.Sprintf("start server err=%v", err))
		}
	}()

	// 等待信号
	<-ctx.Done()

	// 退出服务
	c, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	if err := server.Stop(c); err != nil {
		logrus.Error(err)
	}

	// 清理数据库连接
	cleanup()
}

// 配置文件加载
func loadConfig() (*conf.Bootstrap, error) {
	cfg := new(conf.Bootstrap)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
