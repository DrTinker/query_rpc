package service

import (
	"context"
	"query_rpc/client"
	"query_rpc/grpc_gen/query"
	"query_rpc/models"
	"time"
)

type QueryClientService interface {
	GetQueryByID(ctx context.Context, id int32) (*query.Query, error)
	CreateQuery(ctx context.Context, query *query.Query, id int32) error
}

type queryClientServiceImpl struct {
	client client.DBClient
}

func GetQueryClientService() QueryClientService {
	return &queryClientServiceImpl{
		client: client.GetDBClient(),
	}
}

func (q *queryClientServiceImpl) GetQueryByID(ctx context.Context, id int32) (*query.Query, error) {
	info, err := q.client.GetQueryByID(id)
	if err != nil {
		return nil, err
	}
	query := &query.Query{}
	query.QueryId = info.Query_ID
	query.QueryName = info.Query_Name
	query.State = info.State
	query.Remark = info.Remark
	query.StartTime = info.Start_time.Unix()
	query.EndTime = info.End_time.Unix()
	query.EndMethod = info.End_method
	query.Background = info.Background
	query.Creator = info.Creator
	query.CreateTime = info.Create_time.Unix()
	return query, nil
}

func (q *queryClientServiceImpl) CreateQuery(ctx context.Context, query *query.Query, id int32) error {
	info := &models.Query{}
	info.Query_ID = query.QueryId
	info.Query_Name = query.QueryName
	info.State = query.State
	info.Remark = query.Remark
	info.Start_time = time.Unix(query.StartTime, 0)
	info.End_time = time.Unix(query.EndTime, 0)
	info.End_method = query.EndMethod
	info.Background = query.Background
	info.Creator = query.Creator
	info.Create_time = time.Unix(query.CreateTime, 0)
	err := q.client.CreateQuery(info, id)
	if err != nil {
		return err
	}
	return nil
}
