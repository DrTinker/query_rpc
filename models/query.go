package models

import "time"

type Query struct {
	Query_ID    int64
	Query_Name  string
	State       int32
	Remark      string
	Start_time  time.Time
	End_time    time.Time
	End_method  int32 // 0: 未结束 1：问卷到期 2: 手动终止 ...
	Background  string
	Creator     int64
	Create_time time.Time
	Count       int64
}
