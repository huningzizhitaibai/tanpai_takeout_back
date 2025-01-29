package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"tanpai_takeout_back/config"
	"tanpai_takeout_back/logger"
)

var (
	Config *config.AllConfig //全局Config，用于将所有要用到的配置文件进行存储
	Log    logger.ILog       //用于记录所有的运行日志
	DB     *gorm.DB          //连接的数据库实例
	Redis  *redis.Client     //连接的redis实例

)
