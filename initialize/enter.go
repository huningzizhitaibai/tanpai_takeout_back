package initialize

import (
	"github.com/gin-gonic/gin"
	"tanpai_takeout_back/config"
	"tanpai_takeout_back/global"
	"tanpai_takeout_back/logger"
)

//服务启动时的全局初始化

func GlobalInit() *gin.Engine {
	//配置文件初始化
	global.Config = config.InitLoadConfig()

	//Log初始化
	global.Log = logger.NewLogger(global.Config.Log.Level, global.Config.Log.FilePath)

	//gorm初始化
	global.DB = InitDatabase(global.Config.DataSource.Dsn())

	router := routerInit()
	return router
}
