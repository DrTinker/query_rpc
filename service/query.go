package service

import (
	"context"
	"query_rpc/client"
	"query_rpc/grpc_gen/query"
	"query_rpc/models"
	"time"
)

type QueryClientService interface {
	GetQueryByID(ctx context.Context, id int64) (*query.Query, error)
	CreateQuery(ctx context.Context, query *query.Query, id int64) error
	// 根据用户id批量获取，t为问卷类型
	GetQueryBatch(ctx context.Context, id int64, t int32) ([]*query.Query, error)
}

type queryClientServiceImpl struct {
	client client.DBClient
}

func GetQueryClientService() QueryClientService {
	return &queryClientServiceImpl{
		client: client.GetDBClient(),
	}
}

func (q *queryClientServiceImpl) GetQueryByID(ctx context.Context, id int64) (*query.Query, error) {
	info, err := q.client.GetQueryByID(id)
	if err != nil {
		return nil, err
	}
	query := modelToRPC(info)
	return query, nil
}

func (q *queryClientServiceImpl) CreateQuery(ctx context.Context, query *query.Query, id int64) error {
	info := rpcToModel(query)
	err := q.client.CreateQuery(info, id)
	if err != nil {
		return err
	}
	return nil
}

func (q *queryClientServiceImpl) GetQueryBatch(ctx context.Context, id int64, t int32) ([]*query.Query, error) {
	infos, err := q.client.GetQueryBatch(id, t)
	if err != nil {
		return nil, err
	}
	size := len(infos)
	querys := make([]*query.Query, size)
	for i := 0; i < size; i++ {
		querys[i] = modelToRPC(infos[i])
	}
	return querys, nil
}

func modelToRPC(info *models.Query) *query.Query {
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
	query.Number = info.Count

	return query
}

func rpcToModel(query *query.Query) *models.Query {
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

	return info
}
