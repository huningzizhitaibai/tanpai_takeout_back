package main

import (
	"github.com/gin-gonic/gin"
	"tanpai_takeout_back/global"
	"tanpai_takeout_back/initialize"
)

func main() {

	//初始化所有配置
	router := initialize.GlobalInit()

	//设置运行环境
	gin.SetMode(global.Config.Server.Level)

	router.Run(":" + global.Config.Server.Port)
}
