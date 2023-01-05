// Copyright 2022 Innkeeper zmjaction(zhaomingjun) <748173631@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/zmjaction/zblog

package zblog

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// recommendedHomeDir 定义放置 zblog 服务配置的默认目录.
	recommendedHomeDir = ".zblog"

	// defaultConfigName 指定了 zblog 服务的默认配置文件名.
	defaultConfigName = "zblog.yaml"
)

/*
如果指定了 cfgFile 则直接读取该配置文件，如果没有设置 cfgFile 则在用户主目录中搜索名为
.zblog.yaml 的配置文件，如果找到则读取。如果 cfgFile 为空，并且在用户主目录下没有找到
.zblog.yaml 配置文件，则调用 viper.ReadInConfig() 读取配置文件时报错
*/
// initConfig 设置需要读取的配置文件名、环境变量，并读取配置文件内容到 viper 中.
func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath(".")

		// 设置配置文件格式为 YAML (YAML 格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 ZBLOG，如果是 zblog，将自动转变为大写。
	viper.SetEnvPrefix("ZBLOG")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// 打印 viper 当前使用的配置文件，方便 Debug.
	fmt.Fprintln(os.Stdout, "Using config file:", viper.ConfigFileUsed())
}
