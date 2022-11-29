package oplog_app

import "fmt"

//
//  获取索引名称
//  @Description:
//  @params date     日期 （2022-07）
//  @return string   索引名
//
func getIndexName(index string, date string) string {
	return fmt.Sprint(index, "_", date)
}
