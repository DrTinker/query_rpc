package service

import (
	"context"
	"query_rpc/client"
	"query_rpc/grpc_gen/question"
	"query_rpc/models"
	"time"
)

type ResultClientService interface {
	CommitQuery(ctx context.Context, req *question.CommitQueryReq) error
	GetOptionResultByQuestion(ctx context.Context, id int64) (res *question.OptionData, err error)
	GetBlankResultByQuestion(ctx context.Context, id int64, page int) (res []*question.BlankResult, err error)
}

type resultClientServiceImpl struct {
	client client.DBClient
}

func GetResultClientService() ResultClientService {
	return &resultClientServiceImpl{
		client: client.GetDBClient(),
	}
}

// 保存
func (r *resultClientServiceImpl) CommitQuery(ctx context.Context, req *question.CommitQueryReq) error {
	query := req.GetQueryId()
	user := req.GetUesrId()
	t := time.Unix(req.GetCreateTime(), 0)
	opres := req.GetOpRes()
	blres := req.GetBlRes()

	ops := make([]models.OptionResult, len(opres))
	bls := make([]models.BlankResult, len(blres))

	for i, v := range req.GetOpRes() {
		ops[i] = models.OptionResult{}
		ops[i].Op_Res_ID = v.OpResId
		ops[i].Create_Time = t
		ops[i].Op_ID = v.OptionId
		ops[i].Query_ID = query
		ops[i].Question_ID = v.QuestionId
		ops[i].User = user
	}

	for i, v := range req.GetBlRes() {
		bls[i] = models.BlankResult{}
		bls[i].Bl_Res_ID = v.BlResId
		bls[i].Create_Time = t
		bls[i].Result = v.Result
		bls[i].Query_ID = query
		bls[i].Question_ID = v.QuestionId
		bls[i].User = user
	}
	err := r.client.CreateOptionResultBatch(ops)
	if err != nil {
		return err
	}
	return nil
}

func (r *resultClientServiceImpl) GetOptionResultByQuestion(ctx context.Context, id int64) (res *question.OptionData, err error) {
	data, total, err := r.client.GetOptionResultByQuestion(id)
	infos := make([]*question.OptionData_InnerData, len(data))
	i := 0
	for k, v := range data {
		infos[i] = &question.OptionData_InnerData{}
		infos[i].Count = int64(v)
		infos[i].OpId = k
		i++
	}
	res = &question.OptionData{
		QuestionId: id,
		Total:      total,
		Result:     infos,
	}
	return res, err
}

func (r *resultClientServiceImpl) GetBlankResultByQuestion(ctx context.Context, id int64, page int) (res []*question.BlankResult, err error) {
	data, err := r.client.GetBlankResultByQuestion(id, page)
	res = make([]*question.BlankResult, len(data))
	for i, v := range data {
		res[i] = &question.BlankResult{}
		res[i].CreateTime = v.Create_Time.Unix()
		res[i].Result = v.Result
		res[i].UesrId = v.User
	}
	return res, err
}
