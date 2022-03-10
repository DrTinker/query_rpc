package models

// 用户全信息，用于接口
type UserInfo struct {
	User_name  string  `json:"user_name"`
	User_id    int32   `json:"user_id" gorm:"primaryKey"`
	User_pwd   string  `json:"user_pwd"`
	User_tag   string  `json:"user_tag"` // 兴趣标签的json字符串
	Log        string  `json:"log"`      // 用户备注
	Pass       bool    `json:"pass"`
	Workspace  []int32 `json:"woekspace"` // 工作空间包含问卷ID
	History    []int32 `json:"history"`   // 历史问卷包含问卷ID
	Subscribe  []int32 `json:"subscribe"` // 收藏夹包含问卷ID
	User_phone int     `json:"user_phone"`
	User_email string  `json:"user_email"`
}

// 数据库user表信息，用于gorm
type User struct {
	User_name  string `json:"user_name"`
	User_id    int32  `json:"user_id" gorm:"primaryKey"`
	User_pwd   string `json:"user_pwd"`
	User_tag   string `json:"user_tag"` // 兴趣标签的json字符串
	Log        string `json:"log"`      // 用户备注
	Pass       bool   `json:"pass"`
	User_phone int    `json:"user_phone"`
	User_email string `json:"user_email"`
}

// 数据库space表信息，用于gorm
type Space struct {
	User_id  int32
	Query_id int32
	Type     int16
}

// 数据库查询参数
type SelectParam struct {
	Table  string
	Key    string
	Option string
	Value  interface{}
}
