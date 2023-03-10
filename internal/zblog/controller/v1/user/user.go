// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package user

import (
	"github.com/zmjaction/zblog/internal/zblog/biz"
	"github.com/zmjaction/zblog/internal/zblog/store"
)

// UserController 是 user 模块在 Controller 层的实现，用来处理用户模块的请求.
type UserController struct {
	b biz.IBiz
}

// New 创建一个 user controller.
func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
