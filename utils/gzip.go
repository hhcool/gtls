package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
)

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
