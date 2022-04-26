package controller

import (
	"context"
	"query_rpc/grpc_gen/query"
	"query_rpc/service"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type QueryClientController interface {
	GetQueryByID(ctx context.Context, req *query.GetQueryByIDReq) (*query.Query, error)
	CreateQuery(ctx context.Context, req *query.CreateQueryReq) error
	GetQueryBatch(ctx context.Context, req *query.GetQueryBatchReq) ([]*query.Query, error)
}

type queryClientControllerImpl struct {
	service service.QueryClientService
}

func GetQueryClientController() QueryClientController {
	return &queryClientControllerImpl{
		service: service.GetQueryClientService(),
	}
}

func (q *queryClientControllerImpl) GetQueryByID(ctx context.Context, req *query.GetQueryByIDReq) (*query.Query, error) {
	id := req.GetQueryId()
	if id <= 0 {
		logrus.Errorln("[QueryClientController] GetQueryByID id invaild")
		return nil, errors.Wrap(errors.New("id invaild"), "[QueryClientController] GetQueryByID id err: ")
	}
	info, err := q.service.GetQueryByID(ctx, req.GetQueryId())
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (q *queryClientControllerImpl) CreateQuery(ctx context.Context, req *query.CreateQueryReq) error {
	if req.GetQuery().GetCreator() == 0 {
		req.GetQuery().Creator = req.GetUserId()
	}
	err := q.service.CreateQuery(ctx, req.GetQuery(), req.GetUserId())
	if err != nil {
		return err
	}
	return nil
}

func (q *queryClientControllerImpl) GetQueryBatch(ctx context.Context, req *query.GetQueryBatchReq) ([]*query.Query, error) {
	id := req.GetUserId()
	if id <= 0 {
		logrus.Errorln("[QueryClientController] GetQueryBatch id invaild")
		return nil, errors.Wrap(errors.New("id invaild"), "[QueryClientController] GetQueryBatch id err: ")
	}
	t := req.GetQueryType()
	if t < 0 {
		logrus.Errorln("[QueryClientController] GetQueryBatch type invaild")
		return nil, errors.Wrap(errors.New("type invaild"), "[QueryClientController] GetQueryBatch type err: ")
	}
	querys, err := q.service.GetQueryBatch(ctx, id, t)
	if err != nil {
		return nil, err
	}
	return querys, nil
}
