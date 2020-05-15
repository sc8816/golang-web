package utils

import "encoding/json"

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	Success int = 200
	Fail    int = 401
)

//FormatterResult 返回统一格式
func FormatterResult(code int, msg string, data ...interface{}) []byte {
	var res Result
	if len(data) > 0 {
		res = Result{
			Code: code,
			Msg:  msg,
			Data: data[0],
		}
	} else {
		res = Result{
			Code: code,
			Msg:  msg,
		}
	}
	buf, _ := json.Marshal(res)
	return buf
}
