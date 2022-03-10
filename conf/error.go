package conf

import "github.com/pkg/errors"

// 数据库错误
var DBSelectError = errors.New("DB select error")
var DBInsertError = errors.New("DB insert error")
var DBDeleteError = errors.New("DB delete error")
var DBUpdateError = errors.New("DB update error")

// 数据处理错误
var JsonError = errors.New("JSON parse error")
