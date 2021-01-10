package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()
	linkName := "latest_log.log"
	filePath := "log/log" //日志文件路径
	//scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755) //创建日志文件
	//if err != nil {
	//	fmt.Println("err", err)
	//}
	//
	//logger.Out = scr // 输出日志到文件中

	logger.SetLevel(logrus.DebugLevel) //设置日志级别

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",                  //日志文件名
		retalog.WithMaxAge(7*24*time.Hour),     //最大保存时间
		retalog.WithRotationTime(24*time.Hour), //日志分割时间
		retalog.WithLinkName(linkName),         //建立软连接指向最新日志
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0))) //获取开销时间(毫秒)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    //获取请求状态码
		clientIp := c.ClientIP()           // 获取客户端ip
		userAgent := c.Request.UserAgent() // 获取客户端访问工具
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"IP":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}

	}
}
