package test

import (
	"code-go/util"
	"fmt"
	"testing"
)

func TestRandomCode(t *testing.T) {
	fmt.Println(util.RandomCode(4))
}

func TestRandom(t *testing.T) {
	fmt.Println(util.RandomCodeNumLetter(4))
}
