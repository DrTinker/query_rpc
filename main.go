package main

import (
	"net"
	"query_rpc/client"
	"query_rpc/conf"
	"query_rpc/controller"
	basic "query_rpc/grpc_gen/basic"
	query "query_rpc/grpc_gen/query"
	question "query_rpc/grpc_gen/question"
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

func (q queryService) GetQueryBatch(ctx context.Context, req *query.GetQueryBatchReq) (resp *query.GetQueryBatchResp, err error) {
	c := controller.GetQueryClientController()
	resp = new(query.GetQueryBatchResp)
	querys, err := c.GetQueryBatch(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Querys = querys
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

// question服务
type questionService struct{}

var QuestionService = questionService{}

func (q questionService) SetQuestionBatch(ctx context.Context, req *question.SetQuestionBatchReq) (resp *question.SetQuestionBatchResp, err error) {
	c := controller.GetQuestionClientController()
	resp = new(question.SetQuestionBatchResp)
	err = c.SetQuestionBatch(ctx, req)
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

func (q questionService) GetQuestionSingle(ctx context.Context, req *question.GetQuestionSingleReq) (resp *question.GetQuestionSingleResp, err error) {
	c := controller.GetQuestionClientController()
	resp = new(question.GetQuestionSingleResp)
	question, err := c.GetQuestionSingle(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Question = question
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

func (q questionService) GetQuestionBatch(ctx context.Context, req *question.GetQuestionBatchReq) (resp *question.GetQuestionBatchResp, err error) {
	c := controller.GetQuestionClientController()
	resp = new(question.GetQuestionBatchResp)
	questions, err := c.GetQuestionBatch(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Questions = questions
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

// result服务
type resultService struct{}

var ResultService = resultService{}

func (r resultService) CommitQuery(ctx context.Context, req *question.CommitQueryReq) (resp *question.CommitQueryResp, err error) {
	c := controller.GetResultClientController()
	resp = new(question.CommitQueryResp)
	err = c.CommitQuery(ctx, req)
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

func (r resultService) GetOptionResultByQuestion(ctx context.Context, req *question.GetOptionResultByQuestionReq) (resp *question.GetOptionResultByQuestionResp, err error) {
	c := controller.GetResultClientController()
	resp = new(question.GetOptionResultByQuestionResp)
	data, err := c.GetOptionResultByQuestion(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Data = data
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

func (r resultService) GetBlankResultByQuestion(ctx context.Context, req *question.GetBlankResultByQuestionReq) (resp *question.GetBlankResultByQuestionResp, err error) {
	c := controller.GetResultClientController()
	resp = new(question.GetBlankResultByQuestionResp)
	data, err := c.GetBlankResultByQuestion(ctx, req)
	if err != nil {
		resp.Resp = &basic.RespBody{
			Code:    conf.RPC_ERROR_CODE,
			RespMsg: "Failed!",
		}
		return resp, err
	}
	resp.Result = data
	resp.Resp = &basic.RespBody{
		Code:    conf.RPC_SUCCESS_CODE,
		RespMsg: "Success!",
	}
	return resp, nil
}

func (r resultService) GetResultByQuery(ctx context.Context, req *question.GetResultByQueryReq) (resp *question.GetResultByQueryResp, err error) {
	return nil, nil
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

	// 服务注册
	user.RegisterUserServiceServer(s, UserService)
	query.RegisterQueryServiceServer(s, QueryService)
	question.RegisterQuestionServiceServer(s, QuestionService)
	question.RegisterResultServiceServer(s, ResultService)

	fmt.Println("Listen on " + address)
	s.Serve(listen)
}
