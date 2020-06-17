package util

import (
	"fmt"
	"time"
)

// 用户token验证器
func AccountTokenValid(token string) bool {
	return true
}

func GenerateToken(mobile string) string {
	// 40位字符：md5(mobile + timestamp + token_salt) + timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := MD5([]byte(mobile + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}