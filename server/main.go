package main

import (
	"gin-vue-admin/server/core"
	"gin-vue-admin/server/global"
	"gin-vue-admin/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.ECOVACS_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.ECOVACS_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.ECOVACS_LOG)
	global.ECOVACS_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.ECOVACS_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.ECOVACS_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
