package util

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitSnowflake() {
	snowflake.Epoch = time.Now().UnixNano() / 1e6
	tmp, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	node = tmp
}

func NewID() int64 {
	return node.Generate().Int64()
}
