package ip

import (
	"dtapps/dta/library/service/ip/ip2region"
	v4 "dtapps/dta/library/service/ip/v4"
	v6 "dtapps/dta/library/service/ip/v6"
	"dtapps/dta/library/utils/gopostgresql"
	"os"
	"strings"
)

type App struct {
	V4Region ip2region.Ip2Region
	V4db     v4.Pointer
	V6db     v6.Pointer
	Pgsql    gopostgresql.App // 日志数据库
}

type FileData struct {
	Data []byte
	Path *os.File
}

func (app *App) Ipv4(ip string) (res v4.Result, resInfo ip2region.IpInfo) {
	res = app.V4db.Find(ip)
	resInfo, _ = app.V4Region.MemorySearch(ip)
	// 日志
	go app.postgresqlIpv4Log(res, resInfo)
	return res, resInfo
}

func (app *App) Ipv6(ip string) (res v6.Result) {
	res = app.V6db.Find(ip)
	// 日志
	go app.postgresqlIpv6Log(res)
	return res
}

func (app *App) isIpv4OrIpv6(ip string) string {
	if len(ip) < 7 {
		return ""
	}
	arrIpv4 := strings.Split(ip, ".")
	if len(arrIpv4) == 4 {
		//. 判断IPv4
		for _, val := range arrIpv4 {
			if !app.CheckIpv4(val) {
				return ""
			}
		}
		return ipv4
	}
	arrIpv6 := strings.Split(ip, ":")
	if len(arrIpv6) == 8 {
		// 判断Ipv6
		for _, val := range arrIpv6 {
			if !app.CheckIpv6(val) {
				return "Neither"
			}
		}
		return ipv6
	}
	return ""
}
