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

	//记录信息，存入数据库
	err = sc.service.ShopSignup(ctx, shopSignup)
	if err != nil {
		code = e.ERROR
		ctx.JSON(http.StatusOK, common.Result{
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
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("食品安全证书缺失")
	}

	fileType = util.FileTypeFinder(shopSignup.IDCard1)
	temp = "./source/temp/" + shopSignup.IDCard1
	dest = "./source/shop/idcard1/" + shopSignup.IDNumber + "-" + shopSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("身份证缺失")
	}

	fileType = util.FileTypeFinder(shopSignup.IDCard2)
	temp = "./source/temp/" + shopSignup.IDCard2
	dest = "./source/shop/idcard2/" + shopSignup.IDNumber + "-" + shopSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("店铺经营许可证缺失")
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "success",
	})
}

func (sc *SignupController) DeliverSignup(ctx *gin.Context) {
	code := e.SUCCESS
	DeliverSignup := request.SignupDTO_Deliver{}

	//这里要确保前端传入的表单是不包含文件数据的
	err := ctx.Bind(&DeliverSignup)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("绑定骑手注册信息错误")
		ctx.JSON(http.StatusBadRequest, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	//记录信息，存入数据库
	err = sc.service.DeliverSignup(ctx, DeliverSignup)
	if err != nil {
		code = e.ERROR
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	//将相关的图片从暂存目录移动到相关存储目录中
	//从暂存中找到相应文件
	fileType := util.FileTypeFinder(DeliverSignup.IDCard1)
	temp := "./source/temp/" + DeliverSignup.IDCard1
	dest := "./source/Deliver/idcard1/" + DeliverSignup.IDNumber + "-" + DeliverSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn(err)
	}

	fileType = util.FileTypeFinder(DeliverSignup.IDCard2)
	temp = "./source/temp/" + DeliverSignup.IDCard2
	dest = "./source/Deliver/idcard2/" + DeliverSignup.IDNumber + "-" + DeliverSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn(err)
	}

	if DeliverSignup.IsStudent {
		fileType = util.FileTypeFinder(DeliverSignup.StudentCard)
		temp = "./source/temp/" + DeliverSignup.StudentCard
		dest = "./source/Deliver/studentcard/" + DeliverSignup.IDNumber + "-" + DeliverSignup.Username + fileType
		err = os.Rename(temp, dest)
		if err != nil {
			code = e.ERROR
			global.Log.Warn("学生证缺失")
		}
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "success",
	})
}

func (sc *SignupController) ControllerSignup(ctx *gin.Context) {
	code := e.SUCCESS
	ControllerSignup := request.SignupDTO_Controller{}

	//这里要确保前端传入的表单是不包含文件数据的
	err := ctx.Bind(&ControllerSignup)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("绑定管理员注册信息错误")
		ctx.JSON(http.StatusBadRequest, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	//记录信息，存入数据库
	err = sc.service.ControllerSignup(ctx, ControllerSignup)
	if err != nil {
		code = e.ERROR
		ctx.JSON(http.StatusOK, common.Result{
			Code: code,
			Msg:  err.Error(),
		})
		return
	}

	//将相关的图片从暂存目录移动到相关存储目录中
	//从暂存中找到相应文件

	fileType := util.FileTypeFinder(ControllerSignup.IDCard1)
	temp := "./source/temp/" + ControllerSignup.IDCard1
	dest := "./source/Controller/idcard1/" + ControllerSignup.IDNumber + "-" + ControllerSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("身份证缺失")
	}

	fileType = util.FileTypeFinder(ControllerSignup.IDCard2)
	temp = "./source/temp/" + ControllerSignup.IDCard2
	dest = "./source/Controller/idcard2/" + ControllerSignup.IDNumber + "-" + ControllerSignup.Username + fileType
	err = os.Rename(temp, dest)
	if err != nil {
		code = e.ERROR
		global.Log.Warn("身份证缺失")
	}

	ctx.JSON(http.StatusOK, common.Result{
		Code: code,
		Msg:  "success",
	})
}
