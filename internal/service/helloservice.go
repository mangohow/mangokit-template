package service

import (
	"context"

	errorv1 "github.com/mangohow/mangokit-template/api/errors/v1"
	"github.com/mangohow/mangokit-template/api/helloworld/v1"
	"github.com/mangohow/mangokit-template/internal/dao"
	"github.com/sirupsen/logrus"
)

type GreeterService struct {
	greeterDao *dao.GreeterDao
	logger     *logrus.Logger
}

func NewGreeterService(greeterDao *dao.GreeterDao, logger *logrus.Logger) *GreeterService {
	return &GreeterService{
		greeterDao: greeterDao,
		logger:     logger,
	}
}

func (s *GreeterService) SayHello(ctx context.Context, request *v1.HelloRequest) (*v1.HelloReply, error) {
	s.logger.Info(s.greeterDao.String())

	return &v1.HelloReply{Message: "hello, " + request.Name}, nil
}

func (s *GreeterService) GetError(ctx context.Context, request *v1.GetErrorRequest) (*v1.GetErrorResponse, error) {
	s.logger.Info("GetError ", "req", request)
	return nil, errorv1.ErrorUserNotFound("user not found")
}

func (s *GreeterService) AddUser(ctx context.Context, request *v1.AddUserRequest) (*v1.AddUserResponse, error) {
	s.logger.Info("AddUser ", "req", request)
	return &v1.AddUserResponse{}, nil
}
