// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zmjaction/zblog/internal/pkg/core"
	"github.com/zmjaction/zblog/internal/pkg/errno"
	"github.com/zmjaction/zblog/internal/pkg/log"
	v1 "github.com/zmjaction/zblog/pkg/api/zblog/v1"
)

// 登录 miniblog 并返回一个 JWT Token.
func (ctrl *UserController) Login(c *gin.Context) {
	log.C(c).Infow("Login function called")

	var r v1.LoginRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}
	resp, err := ctrl.b.Users().Login(c, &r)

	ctrl.b.Users()

	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, resp)
}
