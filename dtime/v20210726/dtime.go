package v20210726

import "time"

const (
	RFC822Format  = "Mon, 02 Jan 2006 15:04:05 MST"
	ISO8601Format = "2006-01-02T15:04:05Z"
)

func NowUTCSeconds() int64 { return time.Now().UTC().Unix() }

func NowUTCNanoSeconds() int64 { return time.Now().UTC().UnixNano() }

// GetCurrentDate 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// GetCurrentUnix 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// GetCurrentMilliUnix 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// GetCurrentNanoUnix 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}

// GetCurrentWjDate 获取当前的时间 - 字符串 - 没有间隔
func GetCurrentWjDate() string {
	return time.Now().Format("20060102")
}

func FormatISO8601Date(timestampSecond int64) string {
	tm := time.Unix(timestampSecond, 0).UTC()
	return tm.Format(ISO8601Format)
}
