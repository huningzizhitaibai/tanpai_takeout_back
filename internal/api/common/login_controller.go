package common

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"tanpai_takeout_back/common"
	"tanpai_takeout_back/common/e"
	"tanpai_takeout_back/common/enum"
	"tanpai_takeout_back/common/util"
	"tanpai_takeout_back/global"
	"tanpai_takeout_back/internal/api/request"
	"tanpai_takeout_back/internal/result/commonResult"
	"tanpai_takeout_back/internal/service/commonService"
)

// Controller中只处理数据相关格式，数据是否能够数据是否满足能够认证符合要求在service中进行处理
// 这两个结构体中都有一样的重名函数，但是处理的方法不同
type LoginController struct {
	service commonService.ILoginService
}

// 自定义的一个构造函数
func NewLoginController(service commonService.ILoginService) *LoginController {
	return &LoginController{service: service}
}

// 这个类中的一些相关方法
func (lc *LoginController) CheckUser(ctx *gin.Context) {
	code := e.SUCCESS
	msg := e.Login_OK

	//定义了一个相关前端需要绑定的结构体
	var loginDto request.LoginDTO
	var flag = true

	//进行绑定
	err := ctx.Bind(&loginDto)
	if err != nil {
		flag = false
		global.Log.Debug("param LoginDTO json failed", err.Error())
		return
	}

	//将绑定的结构体直接传入处理函数进行判断处理，同时使用ctx进行并发控制
	userInfo, err := lc.service.CheckUser(ctx, loginDto)

	if err != nil {
		code = e.ERROR
		msg = e.Login_Fal
		flag = false
		global.Log.Debug("CheckUser err", err.Error())
	}

	//将从用户池中查询到的前端需要用到的参数进行返回绑定
	var login_result commonResult.LoginResult
	temp, _ := json.Marshal(userInfo)
	err = json.Unmarshal(temp, &login_result)
	if err != nil {
		flag = false
		global.Log.Debug("json Unmarshal err", err.Error())
	}

	//全部验证没有问题才生成token
	if flag {
		login_result.Token, err = util.GenerateToken(login_result.Username, enum.UserType(login_result.Type), global.Config.Jwt.User.AccessKeySecret)
	} else {
		login_result.Type = -1
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Data: login_result,
		Msg:  msg,
	})

}
