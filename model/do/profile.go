package do

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	QqChat string `gorm:"type:varchar(32)" json:"qq_chat"`
	WeChat string `gorm:"ype:varchar(32)" json:"wechat"`
	Weibo  string `gorm:"type:varchar(32)" json:"weibo"`
	Email  string `gorm:"type:varchar(32)" json:"email"`
	Img    string `gorm:"type:varchar(80)" json:"img"`
	Avatar string `gorm:"type:varchar(80)" json:"avatar"`
}
