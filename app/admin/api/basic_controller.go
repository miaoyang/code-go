package api

import (
	"code-go/common"
	"code-go/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIp(c *gin.Context) {
	address := util.GetIpAddress(c.Request)
	c.JSON(http.StatusOK, common.OkWithData(address))
}
