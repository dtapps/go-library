package ip2region

import (
	_ "embed"
	"errors"
	"go.dtapp.net/gostring"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	IndexBlockLength = 12
)

type Ip2Region struct {
	// db file handler
	dbFileHandler *os.File

	//header block info

	headerSip []int64
	headerPtr []int64
	headerLen int64

	// super block index info
	firstIndexPtr int64
	lastIndexPtr  int64
	totalBlocks   int64

	// for memory mode only
	// the original db binary string

	dbFile string
}

//go:embed ip2region.db
var dbBinStr []byte

type IpInfo struct {
	IP       string `json:"ip,omitempty"`       // 输入的ip地址
	CityID   int64  `json:"city_id,omitempty"`  // 城市ID
	Country  string `json:"country,omitempty"`  // 国家
	Region   string `json:"region,omitempty"`   // 区域
	Province string `json:"province,omitempty"` // 省份
	City     string `json:"city,omitempty"`     // 城市
	ISP      string `json:"isp,omitempty"`      // 运营商
}

func (ip IpInfo) String() string {
	return ip.IP + "|" + strconv.FormatInt(ip.CityID, 10) + "|" + ip.Country + "|" + ip.Region + "|" + ip.Province + "|" + ip.City + "|" + ip.ISP
}

// 获取Ip信息
func getIpInfo(ipStr string, cityId int64, line []byte) (ipInfo IpInfo) {

	lineSlice := strings.Split(string(line), "|")
	length := len(lineSlice)
	ipInfo.CityID = cityId
	if length < 5 {
		for i := 0; i <= 5-length; i++ {
			lineSlice = append(lineSlice, "")
		}
	}

	if lineSlice[0] != "0" {
		ipInfo.Country = gostring.SpaceAndLineBreak(lineSlice[0])
	}
	if lineSlice[1] != "0" {
		ipInfo.Region = gostring.SpaceAndLineBreak(lineSlice[1])
	}
	if lineSlice[2] != "0" {
		ipInfo.Province = gostring.SpaceAndLineBreak(lineSlice[2])
	}
	if lineSlice[3] != "0" {
		ipInfo.City = gostring.SpaceAndLineBreak(lineSlice[3])
	}
	if lineSlice[4] != "0" {
		ipInfo.ISP = gostring.SpaceAndLineBreak(lineSlice[4])
	}

	ipInfo.IP = ipStr
	return ipInfo
}

// MemorySearch memory算法：整个数据库全部载入内存，单次查询都在0.1x毫秒内
func (r *Ip2Region) MemorySearch(ipStr string) (ipInfo IpInfo, err error) {

	ipInfo.IP = ipStr

	if r.totalBlocks == 0 {

		if err != nil {

			return ipInfo, err
		}

		r.firstIndexPtr = getLong(dbBinStr, 0)
		r.lastIndexPtr = getLong(dbBinStr, 4)
		r.totalBlocks = (r.lastIndexPtr-r.firstIndexPtr)/IndexBlockLength + 1
	}

	ip, err := ip2long(ipStr)
	if err != nil {
		return ipInfo, err
	}

	h := r.totalBlocks
	var dataPtr, l int64
	for l <= h {

		m := (l + h) >> 1
		p := r.firstIndexPtr + m*IndexBlockLength
		sip := getLong(dbBinStr, p)
		if ip < sip {
			h = m - 1
		} else {
			eip := getLong(dbBinStr, p+4)
			if ip > eip {
				l = m + 1
			} else {
				dataPtr = getLong(dbBinStr, p+8)
				break
			}
		}
	}
	if dataPtr == 0 {
		return ipInfo, errors.New("not found")
	}

	dataLen := (dataPtr >> 24) & 0xFF
	dataPtr = dataPtr & 0x00FFFFFF
	ipInfo = getIpInfo(ipStr, getLong(dbBinStr, dataPtr), dbBinStr[(dataPtr)+4:dataPtr+dataLen])
	return ipInfo, nil

}

func getLong(b []byte, offset int64) int64 {

	val := int64(b[offset]) |
		int64(b[offset+1])<<8 |
		int64(b[offset+2])<<16 |
		int64(b[offset+3])<<24

	return val

}

func ip2long(IpStr string) (int64, error) {
	bits := strings.Split(IpStr, ".")
	if len(bits) != 4 {
		return 0, errors.New("ip format error")
	}

	var sum int64
	for i, n := range bits {
		bit, _ := strconv.ParseInt(n, 10, 64)
		sum += bit << uint(24-8*i)
	}

	return sum, nil
}

func (r *Ip2Region) OnlineDownload() (err error) {
	tmpData, err := getOnline()
	if err != nil {
		return errors.New("下载失败 %s" + err.Error())
	}
	if err := ioutil.WriteFile("./ip2region.db", tmpData, 0644); err == nil {
		log.Printf("已下载最新 ip2region 数据库 %s ", "./ip2region.db")
	} else {
		return errors.New("保存失败")
	}
	return nil
}
