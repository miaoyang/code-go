package vo

import "code-go/model/do"

type UserLoginReqVo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserResVo struct {
	Username     string     `json:"username"`
	Mobile       string     `json:"mobile"`
	Avatar       string     `json:"avatar"`
	Nickname     *string    `json:"nickname"`
	Introduction *string    `json:"introduction"`
	Status       uint       `json:"status"`
	Roles        []*do.Role `json:"roles"`
}

// ConvertToUserResVo User转换成UserResVo
func ConvertToUserResVo(user *do.User) UserResVo {
	return UserResVo{
		Username:     user.Username,
		Mobile:       user.Mobile,
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		Introduction: user.Introduction,
		Status:       user.Status,
		Roles:        user.Roles,
	}
}
