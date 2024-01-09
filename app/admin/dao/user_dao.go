package dao

import (
	"code-go/core"
	"code-go/model/do"
)

func InsertUser(user do.User) error {
	tx := core.DB.Create(&user)
	if tx.Error != nil {
		core.LOG.Println("insert user in do fail")
		return tx.Error
	}
	return nil
}

func GetUserByUsername(userName string) *do.User {
	var user *do.User
	tx := core.DB.Model(&do.User{}).Where("username=?", userName).Find(&user)
	if tx.Error != nil {
		core.LOG.Println("Query user by username fail")
	}
	return user
}

func GetUser(pageNum, pageSize int) ([]*do.User, int64) {
	var users []*do.User
	var total int64
	DB := core.DB.Model(&do.User{})
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	core.LOG.Println("查询到的所有账号：", len(users))
	if err != nil {
		core.LOG.Println("GetUser fail", err)
		return nil, 0
	}
	err = DB.Count(&total).Error
	if err != nil {
		core.LOG.Println("GetUser count fail", err)
		return nil, 0
	}
	return users, total
}
