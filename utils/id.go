package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"strings"
)

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
