package middleware

import (
	"code-go/app/admin/api"
	_ "code-go/docs"
	"code-go/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// InitRouter 初始化Router
func InitRouter() *gin.Engine {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(util.Logger())
	g.Use(Cors())

	// 需授权
	auth := g.Group("/api")
	{
		auth.GET("/user/getAllUserInfo", api.GetAllUser)
		auth.POST("/user/register", api.Register)
		auth.POST("/user/login", api.Login)
		auth.GET("/user/getUserByName", api.GetUserByUsername)
	}

	// 无需授权
	norm := g.Group("/")
	{
		norm.GET("/getIp", api.GetIp)
		norm.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return g
}
