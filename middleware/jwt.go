package middleware

import (
	"code-go/common"
	"code-go/global"
	"code-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func JwtToken(c *gin.Context) {
	auth := c.Request.Header.Get(global.Authorization)
	if auth == "" {
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.AUTHORIZATION_EMPTY, "Authorization为空"))
		c.Abort()
		return
	}

	checkToken := strings.SplitN(auth, " ", 2)
	if len(checkToken) != 2 || checkToken[0] != global.AuthBearer {
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.AUTHORIZATION_TYPE_WRONG, "Authorization类型错误"))
		c.Abort()
		return
	}
	token, err := util.ParseToken(auth)
	if err != nil {
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.AUTHORIZATION_PARSE_ERROR, "Authorization解析失败"))
		c.Abort()
		return
	}
	if time.Now().Unix() > token.ExpiresAt {
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.AUTHORIZATION_EXPIRETIME, "Authorization过期"))
		c.Abort()
		return
	}

}
