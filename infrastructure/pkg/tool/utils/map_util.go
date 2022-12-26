package utils_tool

import (
	"asm_platform/infrastructure/pkg/slog"
	"go.mongodb.org/mongo-driver/bson"
)

// ConvertFilter map转化为filter
func ConvertFilter(filterData map[string]interface{}) interface{} {
	filter := bson.M{}
	data, err := bson.Marshal(filterData)
	if err != nil {
		slog.Errorf("marshal error: %v", err)
		return filter
	}

	err = bson.Unmarshal(data, filter)
	if err != nil {
		slog.Errorf("unmarshal error: %v", err)
	}

	slog.Infof("filter: %v", filter)

	return filter
}
