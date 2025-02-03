package common

import (
	"github.com/gin-gonic/gin"
	"tanpai_takeout_back/global"
	"tanpai_takeout_back/internal/api/common"
	"tanpai_takeout_back/internal/repository/dao/commonDao"
	"tanpai_takeout_back/internal/service/commonService"
)

type SignupRouter struct {
	service commonService.ISignupService
}

func (sr *SignupRouter) InitApiRouter(parent *gin.RouterGroup) {
	//在注册不同的类型的用户中就直接进行资质的申请
	signupRouter := parent.Group("/signup")
	sr.service = commonService.NewSignupService(
		commonDao.NewSignupDao(global.DB),
	)
	signupCtrl := common.NewSignupController(sr.service)
	{
		signupRouter.POST("/user", signupCtrl.UserSignup)
		signupRouter.POST("/shop", signupCtrl.ShopSignup)
		signupRouter.POST("/deliver", signupCtrl.DeliverSignup)
		signupRouter.POST("/controller", signupCtrl.ControllerSignup)
	}

}
