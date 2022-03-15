package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
)

// Signature
// @Description: sha1签名
// @param params
// @return string
func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
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
