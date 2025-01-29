package initialize

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"tanpai_takeout_back/global"
	"time"
)

var (
	GormToManyRequestError = errors.New("gorm: to many request error")
)

func InitDatabase(dsn string) *gorm.DB {
	var gormLogger logger.Interface
	if gin.Mode() == gin.DebugMode {
		gormLogger = logger.Default.LogMode(logger.Info)
	} else {
		gormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   //data source name
		DefaultStringSize:         256,   //string类型字段的默认长度
		DisableDatetimePrecision:  true,  //禁用datetime精度3
		DontSupportRenameIndex:    true,  //重命名所用采用删除并重新创建的方式
		DontSupportRenameColumn:   true,  //使用change重命名
		SkipInitializeWithVersion: false, //根据版本自动配置
	}), &gorm.Config{
		Logger: gormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	//慢日志中间件
	SlowQueryLog(db)

	//限流器中间件
	//GormRateLimiter(db, rate.NewLimiter(500, 1000))

	return db

}

// SlowQueryLog 慢查询日志
func SlowQueryLog(db *gorm.DB) {
	err := db.Callback().Query().Before("*").Register("slow_query_start", func(d *gorm.DB) {
		now := time.Now()
		d.Set("start_time", now)
	})
	if err != nil {
		panic(err)
	}

	err = db.Callback().Query().After("*").Register("slow_query_end", func(d *gorm.DB) {
		now := time.Now()
		start, ok := d.Get("start_time")
		if ok {
			duration := now.Sub(start.(time.Time))
			// 一般认为 200 Ms 为Sql慢查询
			if duration > time.Millisecond*200 {
				global.Log.Error("慢查询", "SQL:", d.Statement.SQL.String())
			}
		}
	})
	if err != nil {
		panic(err)
	}
}

// GormRateLimiter Gorm限流器 此限流器不能终止GORM查询链。
//func GormRateLimiter(db *gorm.DB, r *rate.Limiter) {
//	err := db.Callback().Query().Before("*").Register("RateLimitGormMiddleware", func(d *gorm.DB) {
//		if !r.Allow() {
//			d.AddError(GormToManyRequestError)
//			global.Log.Error(GormToManyRequestError.Error())
//			return
//		}
//	})
//	if err != nil {
//		panic(err)
//	}
//}
