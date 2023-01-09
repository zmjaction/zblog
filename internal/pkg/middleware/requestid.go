// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zmjaction/zblog/internal/pkg/known"
)

// RequestID 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头中是否有 'X-Request-ID', 如果有则复用，没有则新建
		requestID := c.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将RequestID 保存在 gin.Context 中，方便后边程序使用
		c.Set(known.XRequestIDKey, requestID)
		// 将RequestID保存在HTTP返回头中，header的键为 'X-Request-ID'
		c.Writer.Header().Set(known.XRequestIDKey, requestID)
		// c.Next()：在中间件中调用 Next() 方法，Next() 方法之前的代码会在到达请求方法前执行，Next() 方法之后的代码则在请求方法处理后执行
		c.Next()

	}
}
