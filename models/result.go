package models

import "time"

type BlankResult struct {
	Bl_Res_ID   int64
	Query_ID    int64
	Question_ID int64
	Result      string
	Create_Time time.Time
	User        int64
}

type OptionResult struct {
	Op_Res_ID   int64
	Query_ID    int64
	Question_ID int64
	Op_ID       int64
	Create_Time time.Time
	User        int64
}

type OptionData struct {
	Op_ID int64
	Count int64 // 选项选择人数
}
