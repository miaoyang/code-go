package util

import (
	"code-go/core"
	"github.com/goccy/go-json"
)

// Struct2Json 结构体转为json
func Struct2Json(obj interface{}) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		core.LOG.Println("[Struct2Json] fail ", err)
		return ""
	}
	return string(marshal)
}

// Json2Struct json转为struct
func Json2Struct(str string, obj interface{}) {
	err := json.Unmarshal([]byte(str), obj)
	if err != nil {
		core.LOG.Println("[Json2Struct] fail ", err)
	}
}
