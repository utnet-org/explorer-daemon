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
