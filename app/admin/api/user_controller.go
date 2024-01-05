package api

import (
	"code-go/app/admin/dao"
	"code-go/common"
	"code-go/core"
	"code-go/global"
	"code-go/model/do"
	"code-go/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// Register 用户注册
// @Summary 用户注册
// @Produce json
// @Router /api/user/register [post]
func Register(c *gin.Context) {
	userName := c.Query("username")
	password := c.Query("password")
	if userName == "" || password == "" {
		core.LOG.Println("输入的用户名和密码为空")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.VALILD_FAIL, "输入的用户名或密码为空"))
		return
	}

	// 查询用户
	daoUser := dao.GetUserByUsername(userName)
	if daoUser.Username == userName {
		core.LOG.Printf("用户: %s 已存在\n", userName)
		c.JSON(http.StatusOK, common.OkWithData("输入的用户已存在"))
		return
	}

	// 生成用户
	genPasswd := util.GenPasswd(password)
	var user do.User
	user.Username = userName
	user.Password = genPasswd
	user.Status = core.User_status_OK

	err := dao.InsertUser(user)
	if err != nil {
		core.LOG.Println("插入用户失败 ", err)
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.INSERT_DB_FAIL, "插入用户失败"))
		return
	}
	c.JSON(http.StatusOK, common.Ok())

}

// Login 用户登录
// @Summary 用户登录
// @Produce json
// @Router /api/user/login [post]
func Login(c *gin.Context) {
	userName, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	if userName == "" || password == "" {
		core.LOG.Println("用户登录：输入的用户名和密码为空")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.VALILD_FAIL, "输入的用户名和密码为空"))
		return
	}
	// 校验格式
	isMatched := util.ValidateUserName(userName)
	if !isMatched {
		core.LOG.Println("用户登录：用户名格式校验失败")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.VALILD_FAIL, "用户名格式校验失败"))
		return
	}
	//isMatched = util.ValidatePassword(password)
	//if !isMatched {
	//	global.Log.Println("用户登录：密码格式校验失败")
	//	c.JSON(http.StatusOK, common.FailWithCodeMsg(common.VALILD_FAIL, "密码格式校验失败"))
	//	return
	//}

	// TODO: 生成验证码

	// 数据库查询
	user := dao.GetUserByUsername(userName)
	core.LOG.Println("查询到的用户：", user)
	if user.Username == "" {
		core.LOG.Println("用户登录：未查询到用户")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.USER_NOT_EXIST, common.GetMapInfo(common.USER_NOT_EXIST)))
		return
	}
	//校验登录密码是否和数据库一致
	isPasswordMatch := util.ComparePasswd(user.Password, password)
	if !isPasswordMatch {
		core.LOG.Println("用户登录：用户密码输入错误")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.USER_PASSWORD_NOT_MATCHED,
			common.GetMapInfo(common.USER_PASSWORD_NOT_MATCHED)))
		return
	}
	// TODO：生成token
	token, err := util.GenerateToken(strconv.Itoa(int(user.ID)), user.Username)
	if err != nil {
		core.LOG.Println("用户登录：生成token失败")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.AUTHORIZATION_FAIL, "生成token失败"))
		return
	}
	c.Header(global.Authorization, token)

	// 写入redis
	redisToken := fmt.Sprintf("user-token-%s", userName)
	isOk := core.Redis.SetEX(redisToken, token, 7*24*time.Hour)
	if isOk == false {
		core.LOG.Println("用户登录：设置值到redis")
		c.JSON(http.StatusOK, common.FailWithCodeMsg(common.REDIS_SET_FAIL, "设置值到Redis失败"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func GetUserByUsername(c *gin.Context) {
	userName := c.Query("username")
	user := dao.GetUserByUsername(userName)
	c.JSON(http.StatusOK, common.OkWithData(user))
}

func GetAllUser(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Param("pagenum"))
	pageSize, _ := strconv.Atoi(c.Param("pagesize"))
	if pageNum == 0 {
		pageNum = -1
	}
	if pageSize == 0 {
		pageSize = -1
	}
	users, total := dao.GetUser(pageNum, pageSize)

	c.JSON(http.StatusOK, common.OkWithData(common.NewPageRes(users, total)))
}
