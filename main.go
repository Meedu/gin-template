package main

import (
	"github.com/Meedu/gin-template/core"
	"github.com/Meedu/gin-template/global"
	"github.com/Meedu/gin-template/initialize"
	"go.uber.org/zap"
)

func main() {
	global.MD_VP = core.Viper() // 初始化Viper
	global.MD_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.MD_LOG)
	global.MD_DB = initialize.Gorm() // gorm连接数据库
	core.RunWindowsServer()
}
