package commonRepo

import (
	"context"
	"tanpai_takeout_back/internal/api/request"
)

type SignupRepo interface {
	//User注册
	UserSignup_d(ctx context.Context, userDTO request.SignUpDTO_User) error
	ShopSignup_d(ctx context.Context, shopDTO request.SignUpDTO_Shop) error
	DeliverSignup_d(ctx context.Context, deliverDTO request.SignupDTO_Deliver) error
	ControllerSignup_d(ctx context.Context, controllerDTO request.SignupDTO_Controller) error
}
