package conf

import "time"

const App = "conf/app.ini"

// 数据库
const USER_TABLE = "user"
const Max_Conn = 100
const Max_Idle_Conn = 10
const Max_Idle_Time = time.Second * 30

// user
const User_WorkSpace_Code = 0
const User_History_Code = 1
const User_Subscribe_Code = 2

// db
const Page_Size = 20
