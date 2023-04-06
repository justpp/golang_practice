package errcode

import "fmt"

type Error struct {
	code int
	msg  string
}

var _codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := _codes[code]; ok {
		panic(fmt.Sprintf("错误码:%d 已存在，请换一个", code))
	}
	_codes[code] = msg
	return &Error{code, msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息:%s", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
