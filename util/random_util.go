package util

import (
	"code-go/core"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789"

// GetUuid 生成uuid
func GetUuid() string {
	newUUID, _ := uuid.NewUUID()
	return newUUID.String()
}

// RandomCode 生成只包含数字的随机数
func RandomCode(codeLen int) string {
	if codeLen <= 0 {
		core.LOG.Error("codeLen<=0")
		return ""
	}
	rand.Seed(time.Now().Unix())
	code := make([]byte, codeLen)
	for i := 0; i < codeLen; i++ {
		code[i] = byte(rand.Intn(10) + '0')
	}

	return string(code)
}

// RandomCodeNumLetter 生成包含字母和数字的随机数
func RandomCodeNumLetter(codeLen int) string {
	if codeLen <= 0 {
		core.LOG.Error("codeLen<=0")
		return ""
	}
	code := make([]byte, codeLen)
	for i := 0; i < codeLen; i++ {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}
