package utils

import (
	"encoding/json"
	"github.com/hhcool/gtls/log"
	"reflect"
	"strconv"
	"strings"
)

// NewStructWithDefault
// @Auth: oak  2021-10-15 19:23:01
// @Description:  根据tag=default初始化struct
// @param bean
func NewStructWithDefault(bean interface{}) {
	configType := reflect.TypeOf(bean)
	for i := 0; i < configType.Elem().NumField(); i++ {
		field := configType.Elem().Field(i)
		defaultValue := field.Tag.Get("default")
		if defaultValue == "" {
			continue
		}
		setter := reflect.ValueOf(bean).Elem().Field(i)
		switch field.Type.String() {
		case "int":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "time.Duration":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "string":
			setter.SetString(defaultValue)
		case "bool":
			boolValue, _ := strconv.ParseBool(defaultValue)
			setter.SetBool(boolValue)
		}
	}
}

func StructToString(data interface{}) string {
	if str, err := json.Marshal(data); err != nil {
		log.Errorf("StructToString -> ", err)
		return ""
	} else {
		return string(str)
	}
}

// FirstReal
// @Auth: oak  2021-10-15 18:37:24
// @Description:  获取第一个真值，支持int/string/error/bool
// @param values
// @return interface{}
func FirstReal(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil && value != "" && value != 0 && value != false {
			return value
		}
	}
	return values[:len(values)-1][0]
}

// Int64ArrToString
// @Auth: oak  2021-11-05 00:01:58
// @Description:  int64切片转换成逗号分割的字符串
// @param arr
// @return string
func Int64ArrToString(arr []int64) string {
	var result []string
	for _, item := range arr {
		result = append(result, strconv.FormatInt(item, 10))
	}
	return strings.Join(result, ",")
}
func StringToInt64Arr(str string) []int64 {
	arr := strings.Split(str, ",")
	var result []int64
	for _, item := range arr {
		if i, err := strconv.ParseInt(item, 10, 64); err == nil {
			result = append(result, i)
		}
	}
	return result
}

// SliceChunk
// @Description: 用于将字符串切片分块
// @param src
// @param chunkSize
// @return chunks
func SliceChunk(src []string, chunkSize int) (chunks [][]string) {
	total := len(src)
	chunks = make([][]string, 0)
	if chunkSize < 1 {
		chunkSize = 1
	}
	if total == 0 {
		return
	}

	chunkNum := total / chunkSize
	if total%chunkSize != 0 {
		chunkNum++
	}

	chunks = make([][]string, chunkNum)

	for i := 0; i < chunkNum; i++ {
		for j := 0; j < chunkSize; j++ {
			offset := i*chunkSize + j
			if offset >= total {
				return
			}

			if chunks[i] == nil {
				actualChunkSize := chunkSize
				if i == chunkNum-1 && total%chunkSize != 0 {
					actualChunkSize = total % chunkSize
				}
				chunks[i] = make([]string, actualChunkSize)
			}

			chunks[i][j] = src[offset]
		}
	}
	return
}

// RemoveDuplicateElement
// @Description: 切片数组去重
// @param s
// @return string[]
func RemoveDuplicateElement(s []string) []string {
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
