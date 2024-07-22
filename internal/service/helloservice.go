package service

import (
	"context"
	errorv1 "github.com/mangohow/mangokit-template/api/errors"
	"github.com/mangohow/mangokit-template/api/helloworld"
	"google.golang.org/protobuf/types/known/timestamppb"

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

func (s *GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	s.logger.Info(s.greeterDao.String())

	return &helloworld.HelloReply{Message: "hello, " + request.Name}, nil
}

func (s *GreeterService) GetError(ctx context.Context, request *helloworld.GetErrorRequest) error {
	s.logger.Info("GetError ", "req", request)
	return errorv1.ErrorUserNotFound
}

func (s *GreeterService) AddUser(ctx context.Context, request *helloworld.AddUserRequest) error {
	s.logger.Info("AddUser ", "req", request)
	return nil
}

func (s *GreeterService) GetUserInfo(ctx context.Context, request *helloworld.GetUserInfoRequest) (*helloworld.GetUserInfoResponse, error) {
	return &helloworld.GetUserInfoResponse{
		Id:         request.Id,
		Username:   "test",
		Email:      "test@test.com",
		PhoneNum:   "111",
		CreateTime: timestamppb.Now(),
	}, nil
}
