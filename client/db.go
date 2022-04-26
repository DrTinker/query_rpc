package client

import (
	"query_rpc/models"
	"sync"
)

type DBClient interface {
	// user
	CreateUser(user *models.User) error
	GetUserByUserID(id int64) (*models.User, error)
	GetUserSpace(id int64) ([]models.Space, error)
	// query
	CreateQuery(query *models.Query, id int64) error
	GetQueryByID(id int64) (*models.Query, error)
	GetQueryBatch(id int64, t int32) ([]*models.Query, error)
	// question
	SetOptionQuestionBatch(questions []models.Question, options []models.Option) error // 一个question对应一个[]option，这里的[]option为展开后的数组
	SetBlankQuestionBatch(questions []models.Question, blanks []models.Blank) error
	GetQuestionSingle(id int64) (*models.QuestionInfo, error)
	GetQuestionBatch(id int64) (map[int64]*models.QuestionInfo, error)
	// result
	CreateOptionResultBatch(res []models.OptionResult) error
	CreateBlankResultBatch(res []models.BlankResult) error
	GetOptionResultByQuestion(id int64) (res map[int64]int, total int64, err error)
	GetBlankResultByQuestion(id int64, page int) (res []models.BlankResult, err error)
}

var (
	db     DBClient
	DBOnce sync.Once
)

func GetDBClient() DBClient {
	return db
}

func InitDBClient(client DBClient) {
	DBOnce.Do(
		func() {
			db = client
		},
	)
}
