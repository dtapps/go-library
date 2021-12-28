package ip

import (
	"gopkg.in/dtapps/go-library.v2/service/ip/ip2region"
	v4 "gopkg.in/dtapps/go-library.v2/service/ip/v4"
	v6 "gopkg.in/dtapps/go-library.v2/service/ip/v6"
	"os"
	"strings"
)

type App struct {
	V4Region ip2region.Ip2Region
	V4db     v4.Pointer
	V6db     v6.Pointer
}

type FileData struct {
	Data []byte
	Path *os.File
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
