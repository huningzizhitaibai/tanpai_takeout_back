package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 使用pflag命令行拼接出相应的命令
var envPtr = pflag.String("env", "dev", "Environment: dev or release")

func InitLoadConfig() *AllConfig {
	//使用pflag读取来自，命令行的参数，如果在命令行执行时使用参数默认为dev
	pflag.Parse()

	//使用viper库将存储在配置文件中的数据进行读入
	config := viper.New()
	//设置读取路径名
	config.AddConfigPath("./config")
	//设置读取文件的名字
	config.SetConfigName(fmt.Sprintf("application-%s", *envPtr))
	//设置文档读取类型
	config.SetConfigType("yaml")
	//创建读取载体
	var configData *AllConfig
	//读取配置文件
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Use Viper ReadInConfig Fatal error config err:%s \n", err))
	}

	//将读入的配置文件进行绑定
	err = config.Unmarshal(&configData)
	if err != nil {
		panic(fmt.Errorf("read config file to struct err: %s\n", err))
	}

	//将最后读入的配置数据体进行返回
	return configData
}

// 所有的和配置相关的结构体
type AllConfig struct {
	Server     Server     //服务器地址
	DataSource DataSource //数据库相关
	Redis      Redis
	Log        Log
	Jwt        Jwt    //存储jwt的加密密钥
	AliOss     AliOss //阿里的对象存储，但是目前我没有
	Wechat     Wechat //微信登录相关
}

type Server struct {
	Port  string
	Level string
}

type DataSource struct {
	Host         string
	Port         string
	UserName     string
	Password     string
	DatabaseName string `mapstructure:"db_name"`
	Config       string
}

func (d *DataSource) Dsn() string {
	return d.UserName + ":" + d.Password + "@tcp(" + d.Host + ":" + d.Port + ")/" + d.DatabaseName + "?" + d.Config
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DataBase int `mapstructure:"db_name"`
}

type Log struct {
	Level    string
	FilePath string //日志文件存储位置
}

type Jwt struct {
	Admin JwtOption
	User  JwtOption
}

type JwtOption struct {
	EndPoint        string //设置有效时间
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name"`
}

type AliOss struct {
	EndPoint        string
	AccessKeyId     string `mapstructure:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret"`
	BucketName      string `mapstructure:"bucket_name"`
}

type Wechat struct {
	AppId  string
	Secret string
}
