package util

import "github.com/google/uuid"

// GetUuid 生成uuid
func GetUuid() string {
	newUUID, _ := uuid.NewUUID()
	return newUUID.String()
}
