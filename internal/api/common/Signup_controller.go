package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"tanpai_takeout_back/common"
	"tanpai_takeout_back/common/e"
	"tanpai_takeout_back/common/util"
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

	//将相关的图片从暂存目录移动到相关存储目录中
	//从暂存中找到相应文件
	fileType := util.FileTypeFinder(shopSignup.CertificateForShop)
	temp := "./source/temp/" + shopSignup.CertificateForShop
	dest := "./source/shop/certificateForShop/" + shopSignup.IDNumber + "-" + shopSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn(err)
	}

	fileType = util.FileTypeFinder(shopSignup.CertificateForFood)
	temp = "./source/temp/" + shopSignup.CertificateForFood
	dest = "./source/shop/certificateForFood/" + shopSignup.IDNumber + "-" + shopSignup.Username + fileType
	os.Rename(temp, dest)

	fileType = util.FileTypeFinder(shopSignup.IDCard1)
	temp = "./source/temp/" + shopSignup.IDCard1
	dest = "./source/shop/idcard1/" + shopSignup.IDNumber + "-" + shopSignup.Username + fileType
	os.Rename(temp, dest)

	fileType = util.FileTypeFinder(shopSignup.IDCard2)
	temp = "./source/temp/" + shopSignup.IDCard2
	dest = "./source/shop/idcard2/" + shopSignup.IDNumber + "-" + shopSignup.Username + fileType
	os.Rename(temp, dest)

	//记录信息，存入数据库
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
