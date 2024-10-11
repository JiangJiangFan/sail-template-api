package main

import (
	"sail-chat/bootstrap"
	"sail-chat/global"
	"sail-chat/routers"

	// "sail-chat/cmd"

	"github.com/gin-gonic/gin"
)

func main() {
	// cmd.Execute()
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("sail-chat 启动成功")

	// 初始化数据库
	global.App.DB = bootstrap.InitDB()
	global.App.Log.Info("数据库连接成功")
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			err := db.Close()
			if err != nil {
				return
			}
		}
	}()

	// 初始化gin框架
	// bootstrap.InitializeGin()

	// 初始化中文翻译器
	bootstrap.InitTrans()
	// bootstrap.InitEnglishTranslator()
	// 初始化自定义验证器
	bootstrap.InitValidation()

	r := gin.Default()
	// 启用中间件
	// r.Use(middleware.CrossSite())
	// 添加路由
	routers.CollectRoute(r)
	// 启动服务器
	err := r.Run(":" + global.App.Config.App.Port)
	if err != nil {
		return
	}
}
