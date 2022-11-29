// @Title  日期工具
// @Description:
// @Author: lvyazhou
// @Date: 2022/5/20 10:02

package utils_tool

import (
	constapicode "asm_platform/infrastructure/pkg/constants/api_code"
	"fmt"
	"time"
)

var (
	longFormat      = "2006-01-02 15:04:05"
	shortFormat     = "2006-01-02"
	yearMonthFormat = "2006-01"
)

//
//  GetDateTime
//  @Description: 返回时间段（开始从0点到结束23点59分59秒）
//  @params t
//  @return string
//
func GetDateTime(sDate string, eDate string) (int64, int64) {
	//获取当前时区
	loc, _ := time.LoadLocation("Local")
	//日期当天0点时间戳(拼接字符串)
	startDate := sDate + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := eDate + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//返回当天0点和23点59分的时间戳
	return startTime.UnixMilli(), end.UnixMilli()
}

// GetDateTimeStr 拼接时间
func GetDateTimeStr(sDate string, eDate string) (string, string) {
	//日期当天0点时间戳(拼接字符串)
	startDate := sDate + "_00:00:00"
	//日期当天23时59分时间戳
	endDate := eDate + "_23:59:59"
	//返回当天0点和23点59分的时间戳
	return startDate, endDate
}

//
//  AddTimeByMinutes
//  @Description: 返回毫秒数；时间+计算；
//  @params t
//  @return int64
//
func AddTimeByMinutes(now time.Time, minutes string) int64 {
	mm, _ := time.ParseDuration(minutes)
	mm1 := now.Add(mm)
	return mm1.UnixMilli()
}

//
//  FormatYear
//  @Description: 日期格式化(返回短格式：2006)
//  @params t
//  @return string
//
func FormatYearMonth(t time.Time) string {
	return fmt.Sprint(t.Format(yearMonthFormat))
}

//
//  FormatLong
//  @Description: 日期格式化(返回长格式：2006-01-02 15:04:05)
//  @params t
//  @return strDate
//
func FormatLong(t *time.Time) string {
	if t == nil {
		return ""
	}
	return fmt.Sprint(t.Format(longFormat))
}

//
//  FormatShort
//  @Description: 日期格式化(返回短格式：2006-01-02)
//  @params t
//  @return string
//
func FormatShort(t time.Time) string {
	return fmt.Sprint(t.Format(shortFormat))
}

//
//  FormatShortPlus
//  @Description: 日期格式化(返回短格式：2006-01-02)
//  @params t
//  @return string
//
func FormatShortPlus(t *time.Time) string {
	if t == nil {
		return ""
	}
	return fmt.Sprint(t.Format(shortFormat))
}

//
//  Timestamp2ShortDateStr
//  @Description: 时间戳 转 时间字符串(返回短格式：2006-01-02)
//  @params timestamp
//  @return string
//
func Timestamp2ShortDateStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(shortFormat)
}

//
//  Timestamp2ShortDateStr
//  @Description: 时间戳 转 时间字符串(返回短格式：2006-01-02)
//  @params timestamp
//  @return string
//
func MilTimestamp2ShortDateStr(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(shortFormat)
}

//
//  Timestamp2ShortDateStr
//  @Description: 时间戳 转 时间字符串(返回短格式：2006-01-02 00：00：00)
//  @params timestamp
//  @return string
//
func Timestamp2LongDateStr(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(longFormat)
}

//
//  Timestamp13LongDateStr
//  @Description: 时间戳(13位) 转 时间字符串(返回短格式：2006-01-02 00：00：00)
//  @params timestamp
//  @return string
//
func Timestamp13LongDateStr(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(longFormat)
}

//
//  ShortStrDate2Time
//  @Description: string（2022-01-01） 日期转time.Time
//  @params str
//  @return time.Time
//
func ShortStrDate2Time(str string) (*time.Time, constapicode.SocError) {
	date, err := time.ParseInLocation(shortFormat, str, time.Local)
	if err != nil {
		fmt.Printf("[Format date] date convert fail; function:ShortStrDate2Time(); param:%s, error: %v", str, err)
		return nil, constapicode.DateConvertError
	}
	return &date, constapicode.Success
}

//
//  ShortStrDate2Time
//  @Description: string（2022-01-01） 日期转time.Time
//  @params str
//  @return time.Time
//
func ShortStrDate3Time(str string) (time.Time, constapicode.SocError) {
	date, err := time.ParseInLocation(shortFormat, str, time.Local)
	if err != nil {
		fmt.Printf("[Format date] date convert fail; function:ShortStrDate3Time(); param:%s, error: %v", str, err)
		return time.Now(), constapicode.DateConvertError
	}
	return date, constapicode.Success
}

// LongStrDate2Time 转 time
//  @Description: string（2022-01-01 12:41:58） 日期转time.Time
//  @params str
//  @return time.Time
//
func LongStrDate2Time(str string) (*time.Time, constapicode.SocError) {
	date, err := time.ParseInLocation(longFormat, str, time.Local)
	if err != nil {
		fmt.Errorf("[Format date] date convert fail; function:LongStrDate2Time(); param:%s, error: %v", str, err)
		return nil, constapicode.DateConvertError
	}
	return &date, constapicode.Success
}

//
//  时间戳转time.Time
//  @Description: 时间戳转time.Time
//  @params timestamp
//  @return time.Time
//
func Timestamp2Time(timestamp int64) time.Time {
	unix := time.Unix(timestamp, 0)
	return unix
}

//
//  日期string转时间戳（10位）
//  @Description:
//  @params str   2022-01-01
//  @return int64
//  @return constapicode.SocError
//
func ShortStrDate2Timestamp10(str string) (int64, constapicode.SocError) {
	date2Time, code1 := ShortStrDate2Time(str)
	if code1 != constapicode.Success {
		return 0, code1
	}
	return date2Time.Unix(), constapicode.Success
}

//
//  日期string转时间戳（13位）
//  @Description:
//  @params str   2022-01-01
//  @return int64
//  @return constapicode.SocError
//
func ShortStrDate2Timestamp13(str string) (int64, constapicode.SocError) {
	date2Time, code1 := ShortStrDate2Time(str)
	if code1 != constapicode.Success {
		return 0, code1
	}
	return date2Time.UnixMilli(), constapicode.Success
}

////
////  时间戳转日期string
////  @Description:
////  @params timestamp
////  @return string
////  @return constapicode.SocError
////
//func Timestamp2LongStrDate(timestamp int64) (string, constapicode.SocError) {
//
//
//	return
//}

// FormatTimeToString 格式化 time.Time 类型为字符串，0000-00-00 00:00:00 时返回空
func FormatTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return fmt.Sprint(t.Format(longFormat))
}

// FormatTimePointerToString  指针类型处理
func FormatTimePointerToString(t *time.Time) string {
	if t == nil {
		return ""
	}
	if t.IsZero() {
		return ""
	}
	return fmt.Sprint(t.Format(longFormat))
}

// FormatLongStrToTime string（2022-01-01 12:41:58） 日期转time.Time 转 time
func FormatLongStrToTime(str string) time.Time {
	date, _ := time.ParseInLocation(longFormat, str, time.Local)
	return date
}

//
//  获取传入间隔内所有日期
//  @Description:
//  @params startDate 开始日期 yyyy-mm-dd
//  @params endDate   结束日期 yyyy-mm-dd
//  @return []string  间隔内日期切片
//
func GetBetweenDates(startDate, endDate string) []string {
	if startDate == endDate {
		return []string{startDate}
	}

	d := make([]string, 0)
	timeFormatTpl := longFormat
	if len(timeFormatTpl) != len(startDate) {
		timeFormatTpl = timeFormatTpl[0:len(startDate)]
	}
	date, err := time.Parse(timeFormatTpl, startDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, endDate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = shortFormat
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}
