// Copyright 2022 Innkeeper zmjaction(zhaomingjun) <748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog.

package main

import (
	"os"

	"github.com/zmjaction/zblog/internal/zblog"
	_ "go.uber.org/automaxprocs"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := zblog.NewZBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
