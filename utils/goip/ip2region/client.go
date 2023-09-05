package ip2region

import (
	"errors"
	"github.com/dtapps/go-library/utils/gostring"
	"os"
	"strconv"
	"strings"
)

const (
	IndexBlockLength = 12
)

var dbBuff []byte

type Client struct {
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

func New(filepath string) (*Client, error) {

	var err error
	c := &Client{}

	dbBuff, err = os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func NewBuff(file []byte) (*Client, error) {

	var _ error
	c := &Client{}

	dbBuff = file

	return c, nil
}

// 获取Ip信息
func getIpInfo(ipStr string, cityId int64, line []byte) (result QueryResult) {

	lineSlice := strings.Split(string(line), "|")
	length := len(lineSlice)
	result.CityId = cityId
	if length < 5 {
		for i := 0; i <= 5-length; i++ {
			lineSlice = append(lineSlice, "")
		}
	}

	if lineSlice[0] != "0" {
		result.Country = gostring.SpaceAndLineBreak(lineSlice[0])
	}
	if lineSlice[1] != "0" {
		result.Region = gostring.SpaceAndLineBreak(lineSlice[1])
	}
	if lineSlice[2] != "0" {
		result.Province = gostring.SpaceAndLineBreak(lineSlice[2])
	}
	if lineSlice[3] != "0" {
		result.City = gostring.SpaceAndLineBreak(lineSlice[3])
	}
	if lineSlice[4] != "0" {
		result.Isp = gostring.SpaceAndLineBreak(lineSlice[4])
	}

	result.Ip = ipStr
	return result
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
