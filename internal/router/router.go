package router

import "tanpai_takeout_back/internal/router/common"

type Routers struct {
	//common
	common.LoginRouter
	common.SignupRouter
}

var AllRouter = new(Routers)
