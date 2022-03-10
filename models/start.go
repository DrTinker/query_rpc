package models

import "google.golang.org/grpc"

// query_rpc连接
var RpcConn *grpc.ClientConn

type RpcConfig struct {
	Address                string
	Port                   int
	ClientPoolConnsSizeCap int
	DialTimeout            int
	KeepAlive              int
	KeepAliveTimeout       int
}
