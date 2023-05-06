package main

import (
	"fmt"
	"go.uber.org/zap"
	"shop-mall/global"
	"shop-mall/initilize"
)

func main() {

	// 初始化日志
	initilize.InitLogger()

	// 初始化配置文件
	initilize.InitConfig()

	// 初始化 mysql
	initilize.InitDB()

	// 初始化路由
	Router := initilize.Routers()

	// 初始化翻译
	if err := initilize.InitTrans("zh"); err != nil {
		zap.S().Panic("多语言初始化失败：", err.Error())
	}

	zap.S().Debugf("启动服务器：http://%s:%d/", global.ServerConfig.Host, global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
