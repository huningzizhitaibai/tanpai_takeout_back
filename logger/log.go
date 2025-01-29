package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type ILog interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type LLogger struct {
	logger *logrus.Logger
}

// 用于实现相关接口，在接口中定义方法，需要一个结构体将这些方法进行实现
// 也算是go语言的特点了
type LogEmailHook struct{}

// Levels 需要监控的日志等级，只有命中例表中的日志等级才能触发Hook
func (l *LogEmailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
	}
}

func (l *LogEmailHook) Fire(e *logrus.Entry) error {
	//触发loggerhook函数调用
	fmt.Println("触发loggerHook函数调用")
	return nil
}

func NewLogger(level string, filePath string) ILog {
	//解析日志等级吗？
	parseLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err.Error())
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open log file" + filePath)
		panic(err.Error())
	}
	log := &logrus.Logger{
		Out:   io.MultiWriter(f, os.Stdout),         //设置日志输出路径
		Level: parseLevel,                           //Debug日志等级
		Hooks: make(map[logrus.Level][]logrus.Hook), //初始化Hook Map，
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     false,
		},
	}

	log.AddHook(&LogEmailHook{})
	log.Infof("日志开启成功")
	return &LLogger{logger: log}
}

func (l *LLogger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}
func (l *LLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}
func (l *LLogger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}
func (l *LLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
func (l *LLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}
