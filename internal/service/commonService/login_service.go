package commonService

import (
	"context"
	"tanpai_takeout_back/internal/api/request"
	"tanpai_takeout_back/internal/model"
	"tanpai_takeout_back/internal/repository/commonRepo"
)

//service层定义了对于这个服务应该实现那些功能

type ILoginService interface {
	//处理接口两个参数，分别进行处理并发控制和结构前端传来的数据
	CheckUser(ctx context.Context, dto request.LoginDTO) (model.User, error)
}

// 定义一个结构体进行具体的实现
type LoginServiceImpl struct {
	repo commonRepo.LoginRepo
}

func (s *LoginServiceImpl) CheckUser(ctx context.Context, dto request.LoginDTO) (model.User, error) {
	userInfo, err := s.repo.Check(ctx, model.User{
		Username: dto.Username,
		Password: dto.Password,
	})
	return userInfo, err
}

func NewLoginService(repo commonRepo.LoginRepo) ILoginService {
	return &LoginServiceImpl{repo: repo}
}
