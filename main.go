package main

import (
	"net"
	"query_rpc/client"
	"query_rpc/conf"
	"query_rpc/controller"
	basic "query_rpc/grpc_gen/basic"
	query "query_rpc/grpc_gen/query"
	user "query_rpc/grpc_gen/user"
	_ "query_rpc/start"
	"strconv"

	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// user服务
type userService struct{}

var UserService = userService{}

func (h userService) GetUserByUserID(ctx context.Context, req *user.GetUserByUserIDReq) (*user.GetUserByUserIDResp, error) {
	c := controller.GetUserClientController()
	resp := new(user.GetUserByUserIDResp)
	user, err := c.GetUserByUserID(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.UserInfo = user
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}

	return resp, nil
}

func (h userService) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserResp, error) {
	c := controller.GetUserClientController()
	resp := new(user.CreateUserResp)
	err := c.CreateUser(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}

	return resp, nil
}

func (h userService) GetUserByUserName(ctx context.Context, in *user.GetUserByUserNameReq) (*user.GetUserByUserNameResp, error) {
	return nil, nil
}
func (h userService) DeleteUser(ctx context.Context, in *user.DeleteUserReq) (*user.DeleteUserResp, error) {
	return nil, nil
}

// query服务
type queryService struct{}

var QueryService = queryService{}

func (q queryService) CreateQuery(ctx context.Context, req *query.CreateQueryReq) (resp *query.CreateQueryResp, err error) {
	c := controller.GetQueryClientController()
	resp = new(query.CreateQueryResp)
	err = c.CreateQuery(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

func (q queryService) GetQueryByID(ctx context.Context, req *query.GetQueryByIDReq) (resp *query.GetQueryByIDResp, err error) {
	c := controller.GetQueryClientController()
	resp = new(query.GetQueryByIDResp)
	query, err := c.GetQueryByID(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Query = query
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

func main() {
	cfg, err := client.GetConfigClient().GetRPCConfig()
	if err != nil {
		panic(err)
	}
	address := cfg.Address + ":" + strconv.Itoa(cfg.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	user.RegisterUserServiceServer(s, UserService)
	query.RegisterQueryServiceServer(s, QueryService)

	fmt.Println("Listen on " + address)
	s.Serve(listen)
}
