package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tanpai_takeout_back/internal/router"
)

func routerInit() *gin.Engine {
	r := gin.Default()
	allRouter := router.AllRouter
	//所有路由的一个实例
	//在设计中，这个实例中有一个路由初始化的函数，将所有的路由在当前的这个服务引擎r下进行初始化

	//图片上传接口
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("image")
		//parentPath, _ := os.Getwd()
		finalPath := "./source/temp/" + file.Filename
		err := c.SaveUploadedFile(file, finalPath)
		if err != nil {
			fmt.Println("保存图片失败")
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":      "保存成功",
			"fileName": file.Filename,
		})
	})

	common := r.Group("/common")
	{
		//我自己添加的测试连接的路由
		common.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "connect successfully",
			})
		})

		allRouter.LoginRouter.InitApiRouter(common)
		allRouter.SignupRouter.InitApiRouter(common)

	}

	////shop 路由组，管理与商铺相关的路由
	//shop := r.Group("/shop")
	//{
	//
	//}
	//
	////user 路由组
	//user := r.Group("/user")
	//{
	//
	//}
	//
	////deliver路由组 ，管理骑手相关
	//deliver := r.Group("/deliver")
	//{
	//
	//}
	//
	////controller路由组，管理平台管理员相关
	//common := r.Group("/common")
	//{
	//
	//}
	return r
}
