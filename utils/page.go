package utils

import "reflect"

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
