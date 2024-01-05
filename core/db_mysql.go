package core

import (
	"code-go/model/do"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

// InitMysql 初始化mysql
func InitMysql() (*gorm.DB, error) {
	mysqlConfig := Config.Database.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.Database)
	LOG.Println("dsn: ", dsn)
	dbMysql, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("open db_mysql error ", err)
		return nil, err
	}

	//迁移
	dbMysql.AutoMigrate(&do.User{})

	db, _ := dbMysql.DB()
	//设置连接池的最大闲置连接数
	db.SetMaxIdleConns(10)
	//设置连接池中的最大连接数量
	db.SetMaxOpenConns(100)
	//设置连接的最大复用时间
	db.SetConnMaxLifetime(10 * time.Second)
	DB = dbMysql
	return dbMysql, nil
}
