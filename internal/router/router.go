package router

import "tanpai_takeout_back/internal/router/common"

type Routers struct {
	common.LoginRouter
}

var AllRouter = new(Routers)
