package logger

/* 日志功能：项目请求日志、错误日志、异常日志、程序员日志等
实现程序员日志功能
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 初始化日志
func init() {
	// 设置日志格式
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutPutFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args...)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args...)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args...)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args...)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args...)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args...)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	setOutPutFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args...)
}

func setOutPutFile(level logrus.Level, logName string) {
	if _, err := os.Stat("./runtime/logs"); os.IsNotExist(err) {
		err = os.MkdirAll("./runtime/logs", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' failed: %s", "./runtime/logs", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	logName = path.Join("./runtime/logs", logName+"-"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed: ", err)
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat("./runtime/logs"); os.IsNotExist(err) {
		err = os.MkdirAll("./runtime/logs", 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' failed: %s", "./runtime/logs", err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	logName := path.Join("./runtime/logs", "success-"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
				params.TimeStamp.Format(time.RFC1123),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		},
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat("./runtime/logs"); os.IsNotExist(errDir) {
				errDir = os.MkdirAll("./runtime/logs", 0777)
				if errDir != nil {
					panic(fmt.Errorf("create log dir '%s' failed: %s", "./runtime/logs", errDir))
				}
			}

			timeStr := time.Now().Format("2006-01-02")
			logName := path.Join("./runtime/logs", "error-"+timeStr+".log")

			f, errFile := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if errFile != nil {
				fmt.Println("open log file failed: ", errFile)
			}
			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString(fmt.Sprintf("%s [ERROR] %s\n", timeFileStr, err))
			f.WriteString("stacktrace from panic:" + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%s", err),
			})
			c.Abort()
		}
	}()
	c.Next()
}
