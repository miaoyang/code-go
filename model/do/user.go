package do

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string  `gorm:"type:varchar(20);not null;unique;index:idx_username" json:"username"`
	Password     string  `gorm:"size:255;not null" json:"password,omitempty"`
	Mobile       string  `gorm:"type:varchar(11);not null;unique" json:"mobile"`
	Avatar       string  `gorm:"type:varchar(255)" json:"avatar"`
	Nickname     *string `gorm:"type:varchar(20)" json:"nickname"`
	Introduction *string `gorm:"type:varchar(255)" json:"introduction"`
	Status       uint    `gorm:"type:tinyint(1);default:1;comment:'1正常, 2禁用'" json:"status"`
	Creator      string  `gorm:"type:varchar(20);" json:"creator"`
	Roles        []*Role `gorm:"many2many:user_roles" json:"roles"`
}
