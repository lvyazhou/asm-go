package utils_tool

//
// SubStr
// @Description: 截取字符串
// @params start 起点下标
// @params end 终点下标(不包括)
// @return string
//
func SubStr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}
