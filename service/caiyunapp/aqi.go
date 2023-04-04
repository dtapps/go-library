package caiyunapp

// GetAqi 空气污染
// https://docs.caiyunapp.com/docs/tables/aqi
// https://www.mee.gov.cn/ywgz/fgbz/bz/bzwb/jcffbz/201203/W020120410332725219541.pdf
func GetAqi(aqi float64) string {
	if aqi <= 50 {
		return "优"
	} else if aqi <= 100 {
		return "良"
	} else if aqi <= 150 {
		return "轻度污染"
	} else if aqi <= 200 {
		return "中度污染"
	} else if aqi <= 300 {
		return "重度污染"
	} else if aqi > 300 {
		return "严重污染"
	}
	return "缺数据"
}

// GetAqiColor 空气污染
// https://docs.caiyunapp.com/docs/tables/aqi
// https://www.mee.gov.cn/ywgz/fgbz/bz/bzwb/jcffbz/201203/W020120410332725219541.pdf
func GetAqiColor(aqi float64) string {
	if aqi <= 50 {
		return "green"
	} else if aqi <= 100 {
		return "yellow"
	} else if aqi <= 150 {
		return "orange"
	} else if aqi <= 200 {
		return "red"
	} else if aqi <= 300 {
		return "purple"
	} else if aqi > 300 {
		return "maroon"
	}
	return ""
}
