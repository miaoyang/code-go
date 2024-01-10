package util

import (
	"code-go/core"
	"fmt"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var Log *logrus.Logger

// SetOutputFile 设置输出文件名称，如果没有就创建
func SetOutputFile() (*os.File, string, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, "", err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	filePath := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(filePath); err != nil {
		if _, err := os.Create(filePath); err != nil {
			log.Println(err.Error())
			return nil, "", err
		}
	}
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err: ", err)
		return nil, "", err
	}
	log.Println("create log path: ", filePath)
	return src, filePath, nil
}

func LoggerNorm() *logrus.Logger {
	//建立软连接，需要管理员权限
	linkName := "latest_log.log"
	//设置日志文件的路径
	src, filePath, _ := SetOutputFile()

	//创建日志
	Log = logrus.New()

	Log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	core.LOG = Log

	//输出
	//Log.Out = src
	//Log.SetOutput(src)
	// 同时输出日志到终端和文件中
	multiWriter := io.MultiWriter(os.Stdout, src)
	Log.SetOutput(multiWriter)

	//设置日志级别
	Log.SetLevel(logrus.DebugLevel)

	// 显示日志行数
	//Log.SetReportCaller(true)

	//添加时间分割
	logWriter, _ := retalog.New(
		filePath,
		retalog.WithMaxAge(14*24*time.Hour),    //日志保留时间：2周
		retalog.WithRotationTime(24*time.Hour), //24小时分割一次
		retalog.WithLinkName(linkName),         //建立软连接
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	//实例化
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(Hook)

	return Log
}

//// Logger 日志，此操作可以复用
//func Logger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		startTime := time.Now()
//		c.Next()
//		stopTime := time.Since(startTime)
//		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
//		hostName, err := os.Hostname()
//		if err != nil {
//			hostName = "unknown"
//		}
//		statusCode := c.Writer.Status()
//		clientIp := c.ClientIP()
//		//userAgent := c.Request.UserAgent()
//		dataSize := c.Writer.Size()
//		if dataSize < 0 {
//			dataSize = 0
//		}
//		method := c.Request.Method
//		requestPath := c.Request.RequestURI
//
//		entry := Log.WithFields(logrus.Fields{
//			"HostName":  hostName,
//			"status":    statusCode,
//			"SpendTime": spendTime,
//			"Ip":        clientIp,
//			"Method":    method,
//			"Path":      requestPath,
//			"DataSize":  dataSize,
//			//"Agent":     userAgent, // TODO: UA
//		})
//		if len(c.Errors) > 0 {
//			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
//		}
//		if statusCode >= 500 {
//			entry.Error()
//		} else if statusCode >= 400 {
//			entry.Warn()
//		} else {
//			entry.Info()
//		}
//	}
//}
