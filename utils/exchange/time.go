package exchange

import (
	"time"
)

// 定义北京时区（中国标准时间 CST/GMT+8）
var CstLocation = time.FixedZone("CST", 8*3600)

type SGETimeRes struct {
	MarketClosed bool   `json:"market_closed"`    // 是否已休市
	Reason       string `json:"reason,omitempty"` // 休市原因
}

// 上海黄金交易所（SGE）是中国唯一的黄金交易所
/**
1. 日盘交易时间
上午：09:00 – 11:30
下午：13:30 – 15:30
2. 夜盘交易时间
晚上：20:00 – 次日凌晨 02:30
注意：夜盘属于下一个交易日的开始。例如，周一晚上20:00开始的夜盘，属于周二的交易时段。

3. 周末及节假日
周六、周日 休市
法定节假日（如春节、国庆节等）期间 休市，具体安排以上海黄金交易所公告为准
**/

// SGETime 判断当前时间是否为 SGE 交易时间
func SGETime(now time.Time) (sge SGETimeRes) {
	// 转换为北京时间
	now = now.In(CstLocation)
	weekday := now.Weekday()
	hour, minute := now.Hour(), now.Minute()

	// 周末休市
	if weekday == time.Saturday || weekday == time.Sunday {
		return SGETimeRes{
			MarketClosed: true,
			Reason:       "周末休市",
		}
	}

	// 判断夜盘是否存在
	if weekday == time.Friday && hour >= 20 {
		// 周五晚上没有夜盘
		return SGETimeRes{
			MarketClosed: true,
			Reason:       "周五夜盘休市",
		}
	}

	// 判断是否在日盘时段
	if (hour == 9 && minute >= 0) || (hour > 9 && hour < 11) ||
		(hour == 11 && minute < 30) {
		return SGETimeRes{MarketClosed: false, Reason: "日盘"}
	}

	if (hour == 13 && minute >= 30) || (hour > 13 && hour < 15) ||
		(hour == 15 && minute < 30) {
		return SGETimeRes{MarketClosed: false, Reason: "日盘"}
	}

	// 判断是否在夜盘时段（20:00 - 次日 02:30）
	if hour >= 20 || (hour < 2) || (hour == 2 && minute <= 30) {
		return SGETimeRes{MarketClosed: false, Reason: "夜盘"}
	}

	// 其他时间均休市
	return SGETimeRes{
		MarketClosed: true,
		Reason:       "非交易时段",
	}
}

// SGEActiveTime 判断是否在“活跃时段”
// 即：正常开市时间 ± 30 分钟内 都认为是活跃状态（返回 MarketClosed=false）
func SGEActiveTime(now time.Time, minute int64) SGETimeRes {
	now = now.In(CstLocation)

	if minute == 0 {
		minute = 30
	}

	// 定义当日各交易时间段（含夜盘跨天）
	type timeRange struct {
		start time.Time
		end   time.Time
		label string
	}

	year, month, day := now.Date()
	location := CstLocation

	// 定义时间段
	ranges := []timeRange{
		{time.Date(year, month, day, 9, 0, 0, 0, location), time.Date(year, month, day, 11, 30, 0, 0, location), "日盘"},
		{time.Date(year, month, day, 13, 30, 0, 0, location), time.Date(year, month, day, 15, 30, 0, 0, location), "日盘"},
		{time.Date(year, month, day, 20, 0, 0, 0, location), time.Date(year, month, day+1, 2, 30, 0, 0, location), "夜盘"},
	}

	weekday := now.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return SGETimeRes{MarketClosed: true, Reason: "周末休市"}
	}
	if weekday == time.Friday {
		// 周五夜盘没有
		ranges = ranges[:2]
	}

	for _, tr := range ranges {
		start := tr.start.Add(-time.Duration(minute) * time.Minute) // 提前30分钟
		end := tr.end.Add(time.Duration(minute) * time.Minute)      // 延后30分钟

		if now.After(start) && now.Before(end) {
			return SGETimeRes{MarketClosed: false, Reason: tr.label + "活跃时段"}
		}
	}

	return SGETimeRes{MarketClosed: true, Reason: "休市时段"}
}
