package api

import (
	"code-go/common"
	"code-go/core"
	"code-go/global"
	"code-go/model/vo"
	"code-go/util"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// GetIp 获取请求ip
func GetIp(c *gin.Context) {
	address := util.GetIpAddress(c.Request)
	c.JSON(http.StatusOK, common.OkWithData(address))
}

func GetCaptcha(c *gin.Context) {
	uuid := util.GetUuid()
	code := captcha.NewLen(4)
	core.Redis.SetEX(global.UserCode+uuid, code, 10*time.Minute)
	imgUrl := "/captcha/" + uuid + ".png"
	refresh := imgUrl + "?reload=1"
	//verify := "/captcha/" + captchaId + "/这里替换为正确的验证码进行验证"
	captchaVo := &vo.CaptchaVo{
		CaptchaId:         uuid,
		CaptchaImgUrl:     imgUrl,
		CaptchaRefreshUrl: refresh,
	}
	c.JSON(http.StatusOK, common.OkWithData(captchaVo))
}

func CheckCaptcha(c *gin.Context) {
	captchaId := c.Query("captchaId")
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusOK, common.FailWithMsg("code为空"))
		core.LOG.Debug("code为空")
		return
	}
	ok, redisVal := core.Redis.Get(global.UserCode + captchaId)
	if !ok {
		c.JSON(http.StatusOK, common.FailWithMsg("code缓存为空"))
		core.LOG.Debug("code缓存为空")
		return
	}
	if strings.EqualFold(redisVal, code) {
		c.JSON(http.StatusOK, common.OkWithData("验证成功"))
		return
	} else {
		c.JSON(http.StatusOK, common.FailWithMsg("校验失败"))
	}
}
