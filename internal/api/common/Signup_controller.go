package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tanpai_takeout_back/common"
	"tanpai_takeout_back/common/e"
	"tanpai_takeout_back/global"
	"tanpai_takeout_back/internal/api/request"
	"tanpai_takeout_back/internal/service/commonService"
)

type SignupController struct {
	service commonService.ISignupService
}

func NewSignupController(signService commonService.ISignupService) *SignupController {
	return &SignupController{service: signService}
}

func (sc *SignupController) UserSignup(ctx *gin.Context) {
	code := e.SUCCESS
	userSignup := request.SignUpDTO_User{}
	err := ctx.Bind(&userSignup)
	if err != nil {
		code = e.ERROR
		global.Log.Debug("UserSign 失败")
		return
	}
	err = sc.service.UserSignup(ctx, userSignup)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("SignupController UserSignup Error:", err.Error())
		//注册失败
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "success",
	})
}

func (sc *SignupController) ShopSignup(ctx *gin.Context) {
	code := e.SUCCESS
	shopSignup := request.SignUpDTO_Shop{}

	//这里要确保前端传入的表单是不包含文件数据的
	err := ctx.Bind(&shopSignup)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("绑定商户注册信息错误")
		ctx.JSON(http.StatusBadRequest, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	err = sc.service.ShopSignup(ctx, shopSignup)
	if err != nil {
		code = e.ERROR
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "success",
	})
}
