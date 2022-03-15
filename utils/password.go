package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

// GenerateFromPassword
// @Description: 创建密码
// @param str
// @return string
func GenerateFromPassword(str string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return string(hash)
}

// CompareHashAndPassword
// @Description: 校验密码
// @param pwd1
// @param pwd2
// @return bool
func CompareHashAndPassword(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
