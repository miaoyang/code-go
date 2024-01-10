package dao

import (
	"code-go/core"
	"code-go/model/do"
)

// InsertOperateLog 插入日志数据
func InsertOperateLog(log *do.OperationLog) error {
	err := core.DB.Model(&do.OperationLog{}).Create(log).Error
	if err != nil {
		core.LOG.Println("插入日志失败：", err)
		return err
	}
	return nil
}
