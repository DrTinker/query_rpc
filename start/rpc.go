package start

import (
	"query_rpc/client"
	"query_rpc/models"
	"strconv"

	"google.golang.org/grpc"
)

func initRPC() {
	cfg, err := client.GetConfigClient().GetRPCConfig()
	if err != nil {
		panic(err)
	}
	address := cfg.Address + ":" + strconv.Itoa(cfg.Port)
	models.RpcConn, err = grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}
}
