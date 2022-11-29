package utils_tool

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
)

//node
var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		//fmt.Errorf("初始化ID生成工具失败")
		panic("初始化ID生成工具失败")
	}

	node = n
}

// GenerateUniqueId 生成ID
func GenerateUniqueId() int64 {
	return node.Generate().Int64()
}

// GenerateUUID 生成uuid
func GenerateUUID() string {
	return uuid.New().String()
}

// GenerateUniqueUint64Id 生成ID
func GenerateUniqueUint64Id() uint64 {
	return uint64(node.Generate().Int64())
}
