package common

import (
	"github.com/gin-gonic/gin"
	"tanpai_takeout_back/global"
	"tanpai_takeout_back/internal/api/controller"
	"tanpai_takeout_back/internal/repository/dao/commonDao"
	"tanpai_takeout_back/internal/service/commonService"
)

type LoginRouter struct{}

//在设计上，将登录的接口设计成同一个，节省开发的时间

func (cr *LoginRouter) InitApiRouter(parent *gin.RouterGroup) {

	//创建一个公有的路由组，管理不同的登录接口(不同用户使用不同接口进行登录）
	loginGroup := parent

	//使用DAO层，将与数据库相关的操作进行封装
	//Controller用于将将定义的一些处理函数进行实例化
	//这些函数需要处理与数据库相关的逻辑，所以将Dao层作为依赖注入
	//获得的Ctrl实例可以直接对一些数据库操作进行实现
	loginCtrl := controller.NewLoginController(
		commonService.NewLoginService(commonDao.NewLoginDao(global.DB)),
	)
	{
		loginGroup.POST("/login", loginCtrl.CheckUser)
	}
}
