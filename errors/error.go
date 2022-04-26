package errors

import (
	"encoding/json"
)

const (
	InnError = 101 //内部错误
	DbError = 102 //数据库错误
	ValError = 103 //参数错误

	Success = 200 //内部错误
	Fail = 400 //数据库错误
	Not = 403 //参数错误
)

type Err struct {
	Code int
	Msg   string
}

func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}

func New(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:   msg,
	}
}