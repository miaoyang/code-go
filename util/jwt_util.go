package util

import (
	"code-go/core"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("code-go")

type Claims struct {
	UserId   string
	UserName string
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(userId string, userName string) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claim := &Claims{
		UserId:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "jwt",
			Issuer:    "code-go",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	signedString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*Claims, error) {
	// 解析JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		core.LOG.Println("ParseToken error: ", err)
		return nil, err
	}

	// 验证令牌的有效性
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		core.LOG.Println("ParseToken error: 令牌无效")
		return nil, fmt.Errorf("令牌无效")
	}
}
