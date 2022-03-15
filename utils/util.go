package utils

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/hhcool/gtls/log"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math"
	"math/rand"
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

// Md5
// @Auth: oak  2021-10-15 18:43:27
// @Description:  MD5
// @param v
// @return string
func Md5(v string) string {
	data := []byte(v)
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	return hex.EncodeToString(md5Ctx.Sum(nil))
}

func StructToString(data interface{}) string {
	if str, err := json.Marshal(data); err != nil {
		log.Errorf("StructToString -> ", err)
		return ""
	} else {
		return string(str)
	}
}

// DefaultPage
// @Auth: oak  2021-10-15 18:40:54
// @Description:  从struct中提取page、rows，若不存在则赋予默认值
// @param c
// @return int
// @return int
func DefaultPage(param interface{}) (int, int) {
	immutable := reflect.ValueOf(param)
	page := reflect.Indirect(immutable).FieldByName("Page").Int()
	rows := reflect.Indirect(immutable).FieldByName("Rows").Int()
	if page == 0 {
		page = 1
	}
	if rows == 0 {
		rows = 50
	}
	if rows > 1000 {
		rows = 1000
	}
	return int(page), int(rows)
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

// NewUuid
// @Auth: oak  2021-10-15 18:36:29
// @Description:  UUIDv4
// @return string
func NewUuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

var sn, _ = snowflake.NewNode(1)

// NewSnowflake
// @Auth: oak  2021-10-15 18:36:13
// @Description:  雪花ID
// @return string
func NewSnowflake() string {
	return sn.Generate().String()
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

// RandStr
// @Description: 生成随机数字符串
// @receiver o
// @return string
func RandStr(len int) string {
	return fmt.Sprintf("%0*d", len, rand.Intn(int(math.Pow10(len))))
}

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

// GZip
// @Description: 压缩
// @param str
// @return string
func GZip(str string) string {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, _ = gz.Write([]byte(str))
	_ = gz.Flush()
	_ = gz.Close()
	zipStr := base64.StdEncoding.EncodeToString(b.Bytes())
	return zipStr
}

// UnGZip
// @Description: 解压
// @param str
// @return string
func UnGZip(str string) string {
	data, _ := base64.StdEncoding.DecodeString(str)
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)
	return string(s)
}
