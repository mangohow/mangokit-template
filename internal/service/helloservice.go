package service

import (
	"context"
	"github.com/mangohow/mangokit-template/api/greeter"
	"github.com/mangohow/mangokit-template/internal/dao"
	"github.com/sirupsen/logrus"
)

type GreeterService struct {
	greeterDao *dao.GreeterDao
	logger     *logrus.Logger
}

func (g GreeterService) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloReply, error) {
	if req.Addr == "" {
		return &greeter.HelloReply{Message: "Hello " + req.Name}, nil
	}

	return &greeter.HelloReply{Message: "Nice to meet the friend from " + req.Addr + ". He's name is " + req.Name}, nil
}

func NewGreeterService(greeterDao *dao.GreeterDao, logger *logrus.Logger) *GreeterService {
	return &GreeterService{
		greeterDao: greeterDao,
		logger:     logger,
	}
}
