package controller

import (
	"context"

	"query_rpc/grpc_gen/user"
	"query_rpc/service"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type UserClientController interface {
	GetUserByUserID(ctx context.Context, req *user.GetUserByUserIDReq) (*user.User, error)
	CreateUser(ctx context.Context, req *user.CreateUserReq) error
}

type userClientControllerImpl struct {
	service service.UserClientService
}

func GetUserClientController() UserClientController {
	return &userClientControllerImpl{
		service: service.GetUserClientService(),
	}
}

func (u *userClientControllerImpl) GetUserByUserID(ctx context.Context, req *user.GetUserByUserIDReq) (*user.User, error) {
	if req.GetUserId() <= 0 {
		logrus.Errorln("[UserClientController] GetUserByUserID id invaild")
		return nil, errors.Wrap(errors.New("id invaild"), "[UserClientController] GetUserByUserID id err: ")
	}
	info, err := u.service.GetUserByUserID(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return info, err
}

func (u *userClientControllerImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) error {
	err := u.service.CreateUser(ctx, req.GetUserInfo())
	if err != nil {
		return err
	}
	return nil
}
