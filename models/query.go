package models

import "time"

type Query struct {
	Query_ID    int32
	Query_Name  string
	State       int32
	Remark      string
	Start_time  time.Time
	End_time    time.Time
	End_method  int32 // 0：问卷到期 1: 手动终止 ...
	Background  string
	Creator     int32
	Create_time time.Time
}
