package client

import (
	"query_rpc/models"
	"sync"
)

type DBClient interface {
	// user
	CreateUser(user *models.User) error
	GetUserByUserID(id int32) (*models.User, error)
	GetUserSpace(id int32) ([]models.Space, error)
	// query
	CreateQuery(query *models.Query, id int32) error
	GetQueryByID(id int32) (*models.Query, error)
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
