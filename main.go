package main

import (
	"code-go/core"
	"code-go/docs"
	"code-go/middleware"
	"code-go/util"
	"strconv"
)

//	@title			code-go api
//	@version		1.0
//	@description	code-go项目swagger api介绍
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	猫哥说
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	//cmd.Execute()

	// 加载yaml config配置
	core.InitConfigDev()
	core.PrintConfig()

	// 日志
	util.LoggerNorm()

	// init mysql
	core.InitMysql()

	// init mysql data
	core.InitData()

	// init redis
	core.InitRedis()

	// 路由
	router := middleware.InitRouter()

	// 保存日志到DB，开启多个协程可能存在先后问题TODO:
	for i := 0; i < 1; i++ {
		go middleware.SaveOperationLog()
	}

	// 启动
	port := util.GetAvailablePort(8000, 10000, core.Config.Server.Port)
	core.LOG.Printf("启动的服务：http://localhost:%d", port)
	router.Run(string(":" + strconv.Itoa(port)))

	core.LOG.Println(docs.SwaggerInfo)

}
