package util

import (
	"code-go/core"
	"regexp"
)

// ValidateUserName 校验用户名格式，表示用户名只能包含字母、数字、下划线和减号，长度在4到16之间
func ValidateUserName(username string) bool {
	matched, err := regexp.MatchString(`^[a-zA-Z0-9_-]{4,16}$`, username)
	if err != nil {
		core.LOG.Println("校验用户名失败", err)
		return false
	}
	return matched
}

// ValidatePassword 校验密码格式，表示密码必须同时包含字母和数字，长度至少为8个字符
func ValidatePassword(password string) bool {
	match, err := regexp.MatchString(`^(?=.*[a-zA-Z])(?=.*\d).{8,}$`, password)
	if err != nil {
		core.LOG.Println("校验密码失败", err)
		return false
	}
	return match
}

// ValidatePhone 电话格式校验，以数字开头，长度为11位
func ValidatePhone(phone string) bool {
	reg := regexp.MustCompile(`^[0-9]{11}$`)
	return reg.MatchString(phone)
}
