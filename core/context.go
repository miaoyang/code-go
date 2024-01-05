package core

import (
	"github.com/sirupsen/logrus"
)

const (
	User_status_OK      = 1 // 用户状态
	User_status_Disable = 2
)

//var DB *gorm.DB

var Error error

var LOG *logrus.Logger

//var Config YamlConfig
//
//var Redis *RedisClient
