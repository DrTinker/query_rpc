package controller

import (
	"context"
	"query_rpc/grpc_gen/question"
	"query_rpc/service"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ResultClientController interface {
	CommitQuery(ctx context.Context, req *question.CommitQueryReq) error
	GetOptionResultByQuestion(ctx context.Context, req *question.GetOptionResultByQuestionReq) (res *question.OptionData, err error)
	GetBlankResultByQuestion(ctx context.Context, req *question.GetBlankResultByQuestionReq) (res []*question.BlankResult, err error)
}

type resultClientControllerImpl struct {
	service service.ResultClientService
}

func GetResultClientController() ResultClientController {
	return &resultClientControllerImpl{
		service: service.GetResultClientService(),
	}
}

func (r *resultClientControllerImpl) CommitQuery(ctx context.Context, req *question.CommitQueryReq) error {
	if req.GetQueryId() == 0 {
		logrus.Errorln("[ResultClientController] CommitQuery query_id invaild")
		return errors.Wrap(errors.New("query_id invaild"), "[QueryClientController] CommitQuery query_id err: ")
	}
	if req.GetUesrId() == 0 {
		logrus.Errorln("[ResultClientController] CommitQuery user_id invaild")
		return errors.Wrap(errors.New("query_id invaild"), "[QueryClientController] CommitQuery user_id err: ")
	}
	err := r.service.CommitQuery(ctx, req)
	return err
}

func (r *resultClientControllerImpl) GetOptionResultByQuestion(ctx context.Context, req *question.GetOptionResultByQuestionReq) (res *question.OptionData, err error) {
	if req.GetQuestionId() == 0 {
		logrus.Errorln("[ResultClientController] GetOptionResultByQuestion question_id invaild")
		return nil, errors.Wrap(errors.New("question_id invaild"), "[QueryClientController] GetOptionResultByQuestion question_id err: ")
	}
	res, err = r.service.GetOptionResultByQuestion(ctx, req.GetQuestionId())
	return res, err
}

func (r *resultClientControllerImpl) GetBlankResultByQuestion(ctx context.Context, req *question.GetBlankResultByQuestionReq) (res []*question.BlankResult, err error) {
	if req.GetQuestionId() == 0 {
		logrus.Errorln("[ResultClientController] GetBlankResultByQuestion question_id invaild")
		return nil, errors.Wrap(errors.New("question_id invaild"), "[QueryClientController] GetBlankResultByQuestion question_id err: ")
	}
	if req.GetPageNum() <= 0 {
		logrus.Errorln("[ResultClientController] GetBlankResultByQuestion page  invaild")
		return nil, errors.Wrap(errors.New("page invaild"), "[QueryClientController] GetBlankResultByQuestion page err: ")
	}
	res, err = r.service.GetBlankResultByQuestion(ctx, req.GetQuestionId(), int(req.GetPageNum()))
	return res, err
}
