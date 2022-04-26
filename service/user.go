package service

import (
	"context"
	"encoding/json"
	"query_rpc/client"
	"query_rpc/conf"
	"query_rpc/grpc_gen/user"

	"query_rpc/models"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type UserClientService interface {
	GetUserByUserID(ctx context.Context, id int64) (*user.User, error)
	CreateUser(ctx context.Context, user *user.User) error
}

type userClientServiceImpl struct {
	client client.DBClient
}

func GetUserClientService() UserClientService {
	return &userClientServiceImpl{
		client: client.GetDBClient(),
	}
}

func (u *userClientServiceImpl) CreateUser(ctx context.Context, user *user.User) error {
	info := &models.User{}

	info.User_id = user.UserId
	info.User_name = user.UserName
	info.User_pwd = user.UserPwd
	info.User_email = user.Email
	info.User_phone = int(user.Phone)
	info.Log = user.Log
	info.Pass = user.Pass

	// 解析json
	if user.UserTag != nil {
		tag, err := json.Marshal(user.UserTag)
		if err != nil {
			logrus.Errorln("[UserClientService] CreateUser user_tag json marshal err: %s", err)
			return conf.JsonError
		}
		info.User_tag = string(tag)
	}

	err := u.client.CreateUser(info)
	if err != nil {
		return err
	}
	return nil
}

func (u *userClientServiceImpl) GetUserByUserID(ctx context.Context, id int64) (*user.User, error) {
	info, err := u.client.GetUserByUserID(id)
	if err != nil {
		return nil, err
	}
	spaces, err := u.client.GetUserSpace(id)
	if err != nil {
		return nil, err
	}
	// 解析用户标签
	tag := []string{}
	// 解析json
	if info.User_tag != "" {
		err = json.Unmarshal([]byte(info.User_tag), &tag)
		if err != nil {
			logrus.Errorln("[UserClientService] GetUserByUserID user_tag json unmarshal err: %s", err)
			return nil, conf.JsonError
		}
	}
	// 解析用户空间信息
	work_space := []int64{}
	history := []int64{}
	subscribe := []int64{}
	for i := 0; i < len(spaces); i++ {
		s := spaces[i]
		switch s.Type {
		case conf.User_WorkSpace_Code:
			work_space = append(work_space, s.Query_id)
		case conf.User_History_Code:
			history = append(history, s.Query_id)
		case conf.User_Subscribe_Code:
			subscribe = append(subscribe, s.Query_id)
		default:
			errors.New("[userClientServiceImpl] GetUserByUserID: error user space type")
		}
	}

	res := &user.User{}
	res.UserId = info.User_id
	res.UserName = info.User_name
	res.UserPwd = info.User_pwd
	res.UserTag = tag
	res.Email = info.User_email
	res.Phone = int64(info.User_phone)
	res.Log = info.Log
	res.Pass = info.Pass
	res.Woekspace = work_space
	res.Subscribe = subscribe
	res.History = history

	return res, nil
}
