package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"strings"
	"time"
)

var node *snowflake.Node

func InitSnowflake(startTime string, machineID int64) {
	var st time.Time
	var err error
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		panic(err)
	}
}

// NewSnowflake
// @Auth: oak  2021-10-15 18:36:13
// @Description:  雪花ID
// @return string
func NewSnowflake() string {
	if node == nil {
		panic("Please initialize snowflake first")
	}
	return node.Generate().String()
}

// NewUuid
// @Auth: oak  2021-10-15 18:36:29
// @Description:  UUIDv4
// @return string
func NewUuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
