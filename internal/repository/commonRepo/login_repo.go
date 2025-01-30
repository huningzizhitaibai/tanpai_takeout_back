package commonRepo

import (
	"context"
	"tanpai_takeout_back/internal/model"
)

// 是和login需要，相关的一些参数仓库
type LoginRepo interface {
	//判断用户登录
	Check(ctx context.Context, user model.User_basic) (model.User_basic, error)
}
