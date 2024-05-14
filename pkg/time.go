package pkg

import (
	"strconv"
	"time"
)

func DateFromTimestamp(timestamp int64) string {
	// Convert timestamp to time.Time
	timeValue := time.Unix(timestamp, 0)

	// Format time in a human-readable layout
	dateString := timeValue.Format("2006-01-02 15:04:05")

	// Print the formatted date
	//fmt.Println("Formatted Date:", dateString)
	return dateString
}

func TimeNowUnixStr() string {
	return strconv.FormatInt(TimeNow().Unix(), 10)
}

// 获取中国时区当前时间
func TimeNow() time.Time {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	return time.Now().In(cstSh)
}

// 10位时间戳转为time
func TimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func TimeNanoSecAgo() int64 {
	// 计算24小时前的时间戳（纳秒）
	now := time.Now()
	nanoSec24HoursAgo := now.Add(-24 * time.Hour).UnixNano()
	return nanoSec24HoursAgo
}

func TimeNanoRange(n int64) (int64, int64) {
	//current := time.Unix(0, currentTime)
	//start := current.AddDate(0, 0, -n).Truncate(24 * time.Hour)
	//end := current.Truncate(24 * time.Hour)
	//return start.UnixNano(), end.UnixNano()
	currentTime := time.Now().UnixNano()
	oneDay := int64(24 * time.Hour)
	sevenDays := oneDay * n
	startTime := currentTime - sevenDays
	return startTime, currentTime
}

func TimeMilliRange(n int64) (int64, int64) {
	currentTime := time.Now().UnixMilli()
	oneDay := int64(24 * time.Hour / time.Millisecond)
	sevenDays := oneDay * n
	startTime := currentTime - sevenDays
	return startTime, currentTime
}

// NanoTimestampToDate 将纳秒时间戳转换为日期格式，可以指定精度
func NanoTimestampToDate(nanoTimestamp int64, layout string) string {
	// 计算时间戳的秒部分和纳秒部分
	sec := nanoTimestamp / 1e9
	nsec := nanoTimestamp % 1e9

	// 将秒和纳秒转换为 time.Time
	date := time.Unix(sec, nsec).UTC()

	// 格式化日期
	return date.Format(layout)
}

func MilliTimestampToDate(millis int64, layout string) string {
	// 将毫秒转换为秒和纳秒
	seconds := millis / 1000
	nanoseconds := (millis % 1000) * int64(time.Millisecond)
	// 创建时间对象
	t := time.Unix(seconds, nanoseconds)

	// 格式化日期
	return t.Format(layout)
}

// ConvertNanoToMilli 将纳秒时间戳转换为毫秒时间戳
func ConvertNanoToMilli(nano int64) int64 {
	return nano / 1e6
}

// ConvertMillisToNano 将毫秒时间戳转换为纳秒时间戳
func ConvertMillisToNano(millis int64) int64 {
	return millis * 1e6
}
