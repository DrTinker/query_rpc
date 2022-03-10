package start

import (
	"query_rpc/client"
	"query_rpc/infrastructure/db"
)

func initDB() {
	driver, source, err := client.GetConfigClient().GetDBConfig()
	if err != nil {
		panic(err)
	}
	impl, err := db.NewDBClientImpl(driver, source)
	if err != nil {
		panic(err)
	}
	client.InitDBClient(impl)
}
