package middleware

import (
	"code-go/core"
	"code-go/model/do"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// OperationLogChan 创建缓存日志channel
var OperationLogChan = make(chan *do.OperationLog, 30)

// SaveOperationLog 保存日志到DB
func SaveOperationLog() {
	Logs := make([]do.OperationLog, 0)
	for val := range OperationLogChan {
		Logs = append(Logs, *val)
		core.LOG.Println("通道：", len(Logs), Logs)
		if len(Logs) >= 5 {
			err := core.DB.Create(&Logs).Error
			if err != nil {
				core.LOG.Println("保存日志失败：", err)
			}
			Logs = make([]do.OperationLog, 0)
		}
	}
}

// Logger 日志，此操作可以复用
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		//spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		spendTime := stopTime.Milliseconds()

		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		//userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		requestPath := c.Request.RequestURI

		path := c.FullPath()

		// 获取当前登录用户
		var username string
		ctxUser, exists := c.Get("user")
		if !exists {
			username = "未登录"
		}
		user, ok := ctxUser.(do.User)
		if ok {
			username = user.Username
		} else {
			username = "未登录"
		}

		operationLog := do.OperationLog{
			Username:   username,
			Ip:         c.ClientIP(),
			IpLocation: "",
			Method:     method,
			Path:       path,
			Desc:       "",
			Status:     c.Writer.Status(),
			StartTime:  startTime,
			TimeCost:   spendTime,
			//UserAgent:  c.Request.UserAgent(),
		}
		// 写入通道
		OperationLogChan <- &operationLog

		entry := core.LOG.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      requestPath,
			"DataSize":  dataSize,
			//"Agent":     userAgent, // TODO: UA
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
