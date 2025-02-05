package commonDao

import (
	"context"
	"gorm.io/gorm"
	"tanpai_takeout_back/common/util"
	"tanpai_takeout_back/internal/model"
	"tanpai_takeout_back/internal/repository/commonRepo"
)

type LoginDao struct {
	db *gorm.DB
}

func (c *LoginDao) Check(ctx context.Context, user model.User_basic) (model.User_basic, error) {
	var userInfo model.User_basic
	user.Password = util.PasswordCrypto(user.Password)

	err := c.db.WithContext(ctx).Table("user_basic").Where("username = ? and password = ?", user.Username, user.Password).First(&userInfo).Error
	return userInfo, err
}

func NewLoginDao(db *gorm.DB) commonRepo.LoginRepo {
	return &LoginDao{db: db}
}
