package errcode

import "fmt"

type KError struct {
	Code   int      `json:"code"`
	Msg    string   `json:"msg"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *KError {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在,请更换一个", code))
	}
	codes[code] = msg
	return &KError{Code: code, Msg: msg}
}

func NewErr(msg string) *KError {
	return &KError{Code: 500, Msg: msg}
}

func (e *KError) Error()  string{
	return fmt.Sprintf("错误码:%d, 错误信息: %s", e.Code, e.Msg)
}

