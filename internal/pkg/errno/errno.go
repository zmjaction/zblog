// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package errno

import "fmt"

// Errno 定义了zblog使用的错误类型
type Errno struct {
	HTTP    int
	Code    string
	Message string
}

// Error 实现error 接口红的 `Error` 方法.
func (err *Errno) Error() string {
	return err.Message
}

// SerMessage 设置 Error 类型错误中的Message 字段
func (err *Errno) SerMessage(format string, args ...interface{}) *Errno {
	err.Message = fmt.Sprintf(format, args)
	return err
}

// Decode 尝试从 err 中解析出业务错误码和错误信息.
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message
	default:
	}

	// 默认返回未知错误码和错误信息. 该错误代表服务端出错
	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}
