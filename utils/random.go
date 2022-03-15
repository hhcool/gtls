package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// RandomIntStr
// @Description: 生成随机数字符串
// @receiver o
// @return string
func RandomIntStr(len int) string {
	return fmt.Sprintf("%0*d", len, rand.Intn(int(math.Pow10(len))))
}

// RandomStr
// @Description: 随机生成字符串
// @param length
// @return string
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
