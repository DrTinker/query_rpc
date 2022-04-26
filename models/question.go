package models

// DB
type Question struct {
	Question_ID       int64
	Query_ID          int64
	Question_Name     string
	Question_Describe string
	Mandatory         bool  // 是否必答
	Display           int32 // 展示方式 0：不展示 1：展示 2：作废展示
	Verify            int32 // 验证，仅用于填空题：0：纯数字，1：身份证，2：email，3：手机号
	Number            int32 // 问卷中的题号
	Question_Type     int32 // 0: 单选，1：多选，3：单行填空，4：多行填空
}

type Option struct {
	Op_ID       int64
	Question_ID int64
	Option_Text string
	Option_Name string
	Is_Multiply bool
	Correct     bool
	Relate_IDs  string // 关联全部题目question_id的json字符串
	Jump_ID     int64  // 跳到某个题目的question_id
}

type Blank struct {
	Blank_ID      int64
	Question_ID   int64
	Private_Check bool  // 是否加密展示
	Unique_Check  bool  // 是否全局唯一
	Len_Limit     int32 // 长度限制
}

// 逻辑
type QuestionInfo struct {
	Info   *Question
	Option []Option
	Blank  *Blank
}
