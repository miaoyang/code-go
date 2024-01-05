package test

import (
	"code-go/util"
	"log"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	isMatch := util.ValidatePassword("Root123456")
	log.Println(isMatch)
}
