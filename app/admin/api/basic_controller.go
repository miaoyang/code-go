package api

import (
	"code-go/common"
	"code-go/core"
	"code-go/global"
	"code-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GetIp 获取请求ip
func GetIp(c *gin.Context) {
	address := util.GetIpAddress(c.Request)
	c.JSON(http.StatusOK, common.OkWithData(address))
}

func CheckCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusOK, common.FailWithMsg("code为空"))
		core.LOG.Debug("code为空")
		return
	}
	ok, redisVal := core.Redis.Get(global.UserCode)
	if !ok {
		c.JSON(http.StatusOK, common.FailWithMsg("code缓存为空"))
		core.LOG.Debug("code缓存为空")
		return
	}
	if strings.EqualFold(redisVal, code) {
		c.JSON(http.StatusOK, common.Ok())
		return
	} else {
		c.JSON(http.StatusOK, common.FailWithMsg("校验失败"))
	}
}
