package commonDao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tanpai_takeout_back/internal/api/request"
	"tanpai_takeout_back/internal/repository/commonRepo"
)

type SignupDao struct {
	db *gorm.DB
}

func (s *SignupDao) UserSignup_d(ctx context.Context, user request.SignUpDTO_User) error {
	err := s.db.Table("user").First(&user).Error
	if err != nil {
		err = s.db.WithContext(ctx).Table("user").Create(&user).Error
		return err
	}
	err = errors.New("用户已经存在")
	return err
}

func (s *SignupDao) ShopSignup_d(ctx context.Context, shop request.SignUpDTO_Shop) error {
	err := s.db.Table("shop").First(&shop).Error
	if err != nil {
		err = s.db.WithContext(ctx).Table("shop").Create(&shop).Error
		return err
	}
	err = errors.New("同名商户已经存在")
	return err
}

func (s *SignupDao) DeliverSignup_d(ctx context.Context, deliver request.SignupDTO_Deliver) error {
	err := s.db.WithContext(ctx).Table("deliver").Create(&deliver).Error
	return err
}

func (s *SignupDao) ControllerSignup_d(ctx context.Context, controller request.SignupDTO_Controller) error {
	err := s.db.WithContext(ctx).Table("common").Create(&controller).Error
	return err
}

func NewSignupDao(db *gorm.DB) commonRepo.SignupRepo {
	return &SignupDao{db: db}
}
