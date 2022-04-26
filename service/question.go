package service

import (
	"context"
	"encoding/json"
	"query_rpc/client"
	"query_rpc/conf"
	"query_rpc/grpc_gen/question"
	"query_rpc/models"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type QuestionClientService interface {
	GetQuestionSingle(ctx context.Context, question_id int64) (*question.Question, error)
	GetQuestionBatch(ctx context.Context, query_id int64) ([]*question.Question, error)
	SetQuestionBatch(ctx context.Context, data []*question.Question) error
}

type questionClientServiceImpl struct {
	client client.DBClient
}

func GetQuestionClientService() QuestionClientService {
	return &questionClientServiceImpl{
		client: client.GetDBClient(),
	}
}

func (q *questionClientServiceImpl) GetQuestionSingle(ctx context.Context, question_id int64) (*question.Question, error) {
	info, err := q.client.GetQuestionSingle(question_id)
	if err != nil {
		return nil, err
	}
	res := &question.Question{}
	res.QuestionId = info.Info.Question_ID
	res.QueryId = info.Info.Query_ID
	res.QuestionName = info.Info.Question_Name
	res.QuestionDescribe = info.Info.Question_Describe
	res.QuestionType = info.Info.Question_Type
	res.Number = info.Info.Number
	res.Display = info.Info.Display
	res.Mandatory = info.Info.Mandatory
	res.Verify = info.Info.Verify
	t := info.Info.Question_Type
	if t == 1 || t == 2 {
		ops := make([]*question.Option, len(info.Option))
		for i := 0; i < len(info.Option); i++ {
			ops[i], err = optionInfoToRPC(&info.Option[i])
			if err != nil {
				logrus.Errorln("[questionClientServiceImpl] GetQuestionByID parse option err: %s", err)
				return nil, errors.Wrap(conf.JsonError, "[questionClientServiceImpl] GetQuestionByID err:")
			}
		}
		res.OpInfo = ops
	}
	if t == 3 || t == 4 {
		res.BlInfo = blankInfoToRPC(info.Blank)
	}

	return res, nil
}

func (q *questionClientServiceImpl) GetQuestionBatch(ctx context.Context, query_id int64) ([]*question.Question, error) {
	info, err := q.client.GetQuestionBatch(query_id)
	if err != nil {
		return nil, err
	}
	res := make([]*question.Question, len(info))
	i := 0
	for _, v := range info {
		res[i] = &question.Question{}
		res[i].QuestionId = v.Info.Question_ID
		res[i].QueryId = v.Info.Query_ID
		res[i].QuestionName = v.Info.Question_Name
		res[i].QuestionDescribe = v.Info.Question_Describe
		res[i].QuestionType = v.Info.Question_Type
		res[i].Number = v.Info.Number
		res[i].Display = v.Info.Display
		res[i].Mandatory = v.Info.Mandatory
		res[i].Verify = v.Info.Verify
		t := v.Info.Question_Type
		if t == 1 || t == 2 {
			ops := make([]*question.Option, len(v.Option))
			for i := 0; i < len(v.Option); i++ {
				op, err := optionInfoToRPC(&v.Option[i])
				ops[i] = op
				if err != nil {
					logrus.Errorln("[questionClientServiceImpl] GetQuestions parse option err: %s", err)
					return nil, errors.Wrap(conf.JsonError, "[questionClientServiceImpl] GetQuestions err:")
				}
			}
			res[i].OpInfo = ops
		}
		if t == 3 || t == 4 {
			bl := blankInfoToRPC(v.Blank)
			res[i].BlInfo = bl
		}
		i++
	}
	return res, nil
}

func (q *questionClientServiceImpl) SetQuestionBatch(ctx context.Context, data []*question.Question) error {
	infos := make([]models.Question, len(data))
	bls := make([]models.Blank, len(data))
	var ops []models.Option
	// 统计选择题选项总数
	op_size := 0
	for i, v := range data {
		t := v.QuestionType
		info := questionRPCToInfo(v)
		op_size += len(v.GetOpInfo())
		// 处理选择题
		if t == 1 || t == 2 {
			for _, e := range v.OpInfo {
				op, err := optionRPCToInfo(e)
				if err != nil {
					logrus.Errorln("[questionClientServiceImpl] SetQuestionBatch parse option err: %s", err)
					return errors.Wrap(conf.JsonError, "[questionClientServiceImpl] SetQuestionBatch err:")
				}
				ops = append(ops, *op)
			}
		}
		// 处理填空题
		if t == 3 || t == 4 {
			bl := blankRPCToInfo(v.BlInfo)
			bls[i] = *bl
		}
		infos[i] = *info
	}

	if len(ops) != 0 {
		err := q.client.SetOptionQuestionBatch(infos, ops)
		if err != nil {
			return err
		}
	} else {
		err := q.client.SetBlankQuestionBatch(infos, bls)
		if err != nil {
			return err
		}
	}

	return nil
}

func questionRPCToInfo(data *question.Question) *models.Question {
	info := models.Question{}
	info.Question_ID = data.QuestionId
	info.Query_ID = data.QueryId
	info.Question_Name = data.QuestionName
	info.Question_Describe = data.GetQuestionDescribe()
	info.Question_Type = data.QuestionType
	info.Number = data.Number
	info.Display = data.Display
	info.Mandatory = data.Mandatory
	info.Verify = data.Verify

	return &info
}

func optionRPCToInfo(data *question.Option) (*models.Option, error) {
	op := models.Option{}
	op.Op_ID = data.OpId
	op.Question_ID = data.QuestionId
	op.Option_Name = data.OptionName
	op.Option_Text = data.OptionText
	op.Is_Multiply = data.IsMultipy
	op.Correct = data.Correct
	op.Jump_ID = data.JumpId
	ids, err := json.Marshal(data.RelateIds)
	if err != nil {
		return nil, err
	}
	op.Relate_IDs = string(ids)

	return &op, nil
}

func blankRPCToInfo(data *question.Blank) *models.Blank {
	bl := &models.Blank{}

	bl.Blank_ID = data.BlankId
	bl.Private_Check = data.PrivateCheck
	bl.Question_ID = data.QuestionId
	bl.Unique_Check = data.UniqueCheck
	bl.Len_Limit = data.LenLimit

	return bl
}

func optionInfoToRPC(data *models.Option) (*question.Option, error) {
	target := &question.Option{}
	target.OpId = data.Op_ID
	target.QuestionId = data.Question_ID
	target.OptionName = data.Option_Name
	target.OptionText = data.Option_Text
	target.IsMultipy = data.Is_Multiply
	target.Correct = data.Correct
	target.JumpId = data.Jump_ID
	relates := make([]int64, 0)
	err := json.Unmarshal([]byte(data.Relate_IDs), &relates)
	if err != nil {
		return nil, err
	}
	target.RelateIds = relates
	return target, nil
}

func blankInfoToRPC(data *models.Blank) *question.Blank {
	target := &question.Blank{}
	target.BlankId = data.Blank_ID
	target.QuestionId = data.Question_ID
	target.PrivateCheck = data.Private_Check
	target.UniqueCheck = data.Unique_Check
	target.LenLimit = data.Len_Limit

	return target
}
