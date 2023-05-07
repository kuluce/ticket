package server

import (
	"fmt"
	"os"
	"ticket/backend/internal/config"

	"github.com/toolkits/file"
	"github.com/toolkits/logger"
)

// Start 启动服务器
func Start() {
	logger.Infoln("启动服务")

	// 配置初始化
	if err := config.Read(); err != nil {
		fmt.Printf("无法读取配置文件，err:%v", err)
		if !file.IsExist(config.ConfigFileName) {
			fmt.Printf("配置文件不存在，重新初始化配置文件,请确认配置文件内容后重新运行")
			_ = config.New()
			_ = config.Save()
			os.Exit(1)
		}
	}

	// 日志初始化

	// 数据库初始化

	// 路由初始化

}
