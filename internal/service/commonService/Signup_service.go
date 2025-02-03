package commonService

import (
	"context"
	"errors"
	"tanpai_takeout_back/internal/api/request"
	"tanpai_takeout_back/internal/repository/commonRepo"
)

type ISignupService interface {
	UserSignup(context.Context, request.SignUpDTO_User) error
	ShopSignup(context.Context, request.SignUpDTO_Shop) error
	DeliverSignup(ctx context.Context, deliver request.SignupDTO_Deliver) error
	ControllerSignup(ctx context.Context, deliver request.SignupDTO_Controller) error
}
type SignupImpl struct {
	repo commonRepo.SignupRepo
}

// 普通用户注册
func (si *SignupImpl) UserSignup(ctx context.Context, user request.SignUpDTO_User) error {
	err := si.repo.UserSignup_d(ctx, user)
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func (si *SignupImpl) ShopSignup(ctx context.Context, user request.SignUpDTO_Shop) error {
	err := si.repo.ShopSignup_d(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (si *SignupImpl) DeliverSignup(ctx context.Context, deliver request.SignupDTO_Deliver) error {
	err := si.repo.DeliverSignup_d(ctx, deliver)
	if err != nil {
		return err
	}
	return nil
}
func (si *SignupImpl) ControllerSignup(ctx context.Context, deliver request.SignupDTO_Controller) error {
	err := si.repo.ControllerSignup_d(ctx, deliver)
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func NewSignupService(repo commonRepo.SignupRepo) ISignupService {
	return &SignupImpl{repo: repo}
}
