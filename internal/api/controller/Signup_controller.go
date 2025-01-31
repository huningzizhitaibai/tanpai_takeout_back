package controller

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
