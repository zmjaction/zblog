// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package store

import (
	"context"

	"gorm.io/gorm"

	"github.com/zmjaction/zblog/internal/pkg/model"
)

// UserStore 定义了 user 模块在 store 层所实现的方法
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

// UserStore 接口的实现
type users struct {
	db *gorm.DB
}

// 确保 users 实现了UserStore 接口
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

// Create 创建一条 user数据
func (u users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
