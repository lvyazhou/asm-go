// @Title  json转map
// @Description:
// @Author: lvyazhou
// @Date: 2022/6/16 14:39

package utils_tool

import (
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"asm_platform/infrastructure/pkg/slog"
	"encoding/json"
)

// MapKeyExtract
//  提取map指定key的value
//  @Description:
//  @params key      map key
//  @params mapData  map数据
//  @return map[string]interface{}
//  @return constapicode.SocError
//
func MapKeyExtract(key string, mapData map[string]interface{}) (map[string]interface{}, constapicode.SocError) {
	value, err := json.Marshal(mapData[key])
	if err != nil {
		slog.Errorf("json to map convert fail,err: %v", err)
		return nil, constapicode.ErrorOnShallowCopy
	}
	valueMap, convertCode := Json2Map(string(value))
	if convertCode != constapicode.Success {
		return nil, convertCode
	}
	return valueMap, constapicode.Success
}

// Json2Map
//  string转json
//  @Description:
//  @params jsonStr
//  @return map[string]interface{}
//  @return constapicode.SocError
//
func Json2Map(jsonStr string) (map[string]interface{}, constapicode.SocError) {
	maps := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &maps)
	if err != nil {
		slog.Errorf("json to map convert fail,err: %v", err)
		return nil, constapicode.ErrorOnShallowCopy
	}
	return maps, constapicode.Success
}

// Json2SliceMap
//  json转切片map
//  @Description:
//  @params jsonStr    stringJson
//  @return []map[string]string
//  @return constapicode.SocError
//
func Json2SliceMap(jsonStr string) ([]map[string]string, constapicode.SocError) {
	maps := make([]map[string]string, 0)
	err := json.Unmarshal([]byte(jsonStr), &maps)
	if err != nil {
		slog.Errorf("json to []map convert fail,err: %v", err)
		return nil, constapicode.ErrorOnShallowCopy
	}
	return maps, constapicode.Success
}

// Json2Struct
//  json转struct
//  @Description:
//  @params jsonStr    stringJson
//  @return interface{}
//  @return constapicode.SocError
//
func Json2Struct(jsonStr string) (interface{}, constapicode.SocError) {
	var inf interface{}
	err := json.Unmarshal([]byte(jsonStr), &inf)
	if err != nil {
		slog.Errorf("json to interface convert fail,err: %v", err)
		return nil, constapicode.ErrorOnShallowCopy
	}
	return inf, constapicode.Success
}

// Struct2Json
//  json转struct
//  @Description:
//  @params inf    interface
//  @return string
//  @return constapicode.SocError
//
func Struct2Json(inf interface{}) (string, constapicode.SocError) {
	data, err := json.Marshal(inf)
	if err != nil {
		slog.Errorf("interface to json convert fail,err: %v", err)
		return "", constapicode.ErrorOnShallowCopy
	}
	return string(data), constapicode.Success
}
