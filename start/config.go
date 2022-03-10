package start

import (
	"query_rpc/client"
	"query_rpc/conf"
	"query_rpc/infrastructure/config"
)

func initConfig() {
	impl := config.NewConfigClientImpl()
	err := impl.Load(conf.App)
	if err != nil {
		panic(err)
	}

	client.InitConfigClient(impl)
}
