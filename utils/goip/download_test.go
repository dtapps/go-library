package goip

import (
	"github.com/dtapps/go-library/utils/goip/geoip"
	"github.com/dtapps/go-library/utils/goip/ip2region"
	"github.com/dtapps/go-library/utils/goip/ip2region_v2"
	"github.com/dtapps/go-library/utils/goip/ipv6wry"
	"github.com/dtapps/go-library/utils/goip/qqwry"
	"testing"
)

func TestV4OnlineDownload(t *testing.T) {
	qqwry.OnlineDownload()
}

func TestV6OnlineDownload(t *testing.T) {
	ipv6wry.OnlineDownload()
}

func TestIp2regionOnlineDownload(t *testing.T) {
	ip2region.OnlineDownload()
}

func TestIp2regionV2OnlineDownload(t *testing.T) {
	ip2region_v2.OnlineDownload()
}

func TestGeoIpOnlineDownload(t *testing.T) {
	geoip.OnlineDownload(geoip.GetGeoLite2CountryDownloadUrl("bb26plSFSVqDCJen"), "GeoLite2-Country_20230602.tar.gz")
	geoip.OnlineDownload(geoip.GetGeoLite2CityDownloadUrl("bb26plSFSVqDCJen"), "GeoLite2-City_20230602.tar.gz")
	geoip.OnlineDownload(geoip.GetGeoLite2AsnDownloadUrl("bb26plSFSVqDCJen"), "GeoLite2-ASN_20230602.tar.gz")
}
