package client

import (
	"query_rpc/models"
	"sync"
)

type ConfigClient interface {
	Load(path string) error
	GetDBConfig() (driver, source string, err error)
	GetRPCConfig() (*models.RpcConfig, error)
}

var (
	configClient ConfigClient
	configOnce   sync.Once
)

func GetConfigClient() ConfigClient {
	return configClient
}

func InitConfigClient(client ConfigClient) {
	configOnce.Do(
		func() {
			configClient = client
		},
	)
}
