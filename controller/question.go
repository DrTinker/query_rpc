package controller

import (
	"context"
	"query_rpc/conf"
	"query_rpc/grpc_gen/question"
	"query_rpc/service"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type QuestionClientController interface {
	GetQuestionSingle(ctx context.Context, req *question.GetQuestionSingleReq) (*question.Question, error)
	GetQuestionBatch(ctx context.Context, req *question.GetQuestionBatchReq) ([]*question.Question, error)
	SetQuestionBatch(ctx context.Context, req *question.SetQuestionBatchReq) error
}

type questionClientControllerImpl struct {
	service service.QuestionClientService
}

func GetQuestionClientController() QuestionClientController {
	return &questionClientControllerImpl{
		service: service.GetQuestionClientService(),
	}
}

func (q *questionClientControllerImpl) GetQuestionSingle(ctx context.Context, req *question.GetQuestionSingleReq) (*question.Question, error) {
	if req.GetQuestionId() == 0 {
		logrus.Errorln("[questionClientControllerImpl] GetQuestionSingle question id empty")
		return nil, errors.Wrap(conf.ParamError, "[questionClientControllerImpl] GetQuestionSingle question id empty:")
	}
	data, err := q.service.GetQuestionSingle(ctx, req.GetQuestionId())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (q *questionClientControllerImpl) GetQuestionBatch(ctx context.Context, req *question.GetQuestionBatchReq) ([]*question.Question, error) {
	if req.GetQueryId() == 0 {
		logrus.Errorln("[questionClientControllerImpl] GetQuestionSingle query id empty")
		return nil, errors.Wrap(conf.ParamError, "[questionClientControllerImpl] GetQuestionSingle query id empty:")
	}
	data, err := q.service.GetQuestionBatch(ctx, req.GetQueryId())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (q *questionClientControllerImpl) SetQuestionBatch(ctx context.Context, req *question.SetQuestionBatchReq) error {
	err := q.service.SetQuestionBatch(ctx, req.GetQuestion())
	if err != nil {
		return err
	}
	return nil
}
