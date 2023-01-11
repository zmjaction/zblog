// Copyright 2022 Innkeeper zmjaction(zhaomingjun) &lt;748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	// 全局变量，已经初始化好的S实例
	S *datastore
)

// IStore 定义类 Store 层需要实现的方法
type IStore interface {
	Users() UserStore
}

// datastore 是 IStore 的一个具体实现.
type datastore struct {
	db *gorm.DB
}

// 确保 datastore 实现了 IStore 接口.
var _ IStore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db: db}
	})
	return S
}

// Users 返回一个实现了 UserStore 接口的实例.
func (ds *datastore) Users() UserStore {
	return newUsers(ds.db)
}
