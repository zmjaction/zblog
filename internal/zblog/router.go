// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package zblog

import (
	"github.com/gin-gonic/gin"
	mw "github.com/zmjaction/zblog/internal/pkg/middleware"

	"github.com/zmjaction/zblog/internal/pkg/core"
	"github.com/zmjaction/zblog/internal/pkg/errno"
	"github.com/zmjaction/zblog/internal/pkg/log"
	"github.com/zmjaction/zblog/internal/zblog/controller/v1/user"
	"github.com/zmjaction/zblog/internal/zblog/store"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
			userv1.PUT(":name/change-password", uc.ChangePassword)
			userv1.Use(mw.Authn())
		}
	}

	return nil
}
