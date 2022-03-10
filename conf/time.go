package conf

import "time"

// RPC连接相关参数
const DefaultClientPoolConnsSizeCap = 5
const DefaultDialTimeout = 5 * time.Second
const DefaultKeepAlive = 30 * time.Second
const DefaultKeepAliveTimeout = 10 * time.Second
