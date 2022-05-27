package v6

import (
	_ "embed"
	"encoding/binary"
	"errors"
	"go.dtapp.net/gostring"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"strings"
)

var (
	header  []byte
	country []byte
	area    []byte
	v6ip    uint64
	offset  uint32
	start   uint32
	end     uint32
)

type Result struct {
	IP       string `json:"ip,omitempty"`       // 输入的ip地址
	Country  string `json:"country,omitempty"`  // 国家
	Province string `json:"province,omitempty"` // 省份
	City     string `json:"city,omitempty"`     // 城市
	Area     string `json:"area,omitempty"`     // 区域
	Isp      string `json:"isp,omitempty"`      // 运营商
}

//go:embed ipv6wry.db
var dat []byte

type Pointer struct {
	Offset   uint32
	ItemLen  uint32
	IndexLen uint32
}

// InitIPV4Data 加载
func (q *Pointer) InitIPV4Data() int64 {
	buf := dat[0:8]
	start := binary.LittleEndian.Uint32(buf[:4])
	end := binary.LittleEndian.Uint32(buf[4:])

	return int64((end-start)/7 + 1)
}

// ReadData 从文件中读取数据
func (q *Pointer) readData(length uint32) (rs []byte) {
	end := q.Offset + length
	dataNum := uint32(len(dat))
	if q.Offset > dataNum {
		return nil
	}

	if end > dataNum {
		end = dataNum
	}
	rs = dat[q.Offset:end]
	q.Offset = end
	return rs
}

// Find ip地址查询对应归属地信息
func (q *Pointer) Find(ip string) (res Result) {

	res = Result{}
	res.IP = ip
	q.Offset = 0

	tp := big.NewInt(0)
	op := big.NewInt(0)
	tp.SetBytes(net.ParseIP(ip).To16())
	op.SetString("18446744073709551616", 10)
	op.Div(tp, op)
	tp.SetString("FFFFFFFFFFFFFFFF", 16)
	op.And(op, tp)

	v6ip = op.Uint64()
	offset = q.searchIndex(v6ip)
	q.Offset = offset

	country, area = q.getAddr()

	// 解析地区数据
	info := strings.Split(string(country), "\t")
	if len(info) > 0 {
		i := 1
		for {
			if i > len(info) {
				break
			}
			switch i {
			case 1:
				res.Country = info[i-1]
				res.Country = gostring.SpaceAndLineBreak(res.Country)
			case 2:
				res.Province = info[i-1]
				res.Province = gostring.SpaceAndLineBreak(res.Province)
			case 3:
				res.City = info[i-1]
				res.City = gostring.SpaceAndLineBreak(res.City)
			case 4:
				res.Area = info[i-1]
				res.Area = gostring.SpaceAndLineBreak(res.Area)
			}
			i++ // 自增
		}
	} else {
		res.Country = string(country)
		res.Country = gostring.SpaceAndLineBreak(res.Country)
	}
	// 运营商
	res.Isp = string(area)

	// Delete ZX (防止不相关的信息产生干扰）
	if res.Isp == "ZX" || res.Isp == "" {
		res.Isp = ""
	} else {
		res.Isp = " " + res.Isp
	}

	res.Isp = gostring.SpaceAndLineBreak(res.Isp)

	return
}

func (q *Pointer) getAddr() ([]byte, []byte) {
	mode := q.readData(1)[0]
	if mode == 0x01 {
		// [IP][0x01][国家和地区信息的绝对偏移地址]
		q.Offset = byteToUInt32(q.readData(3))
		return q.getAddr()
	}
	// [IP][0x02][信息的绝对偏移][...] or [IP][国家][...]
	_offset := q.Offset - 1
	c1 := q.readArea(_offset)
	if mode == 0x02 {
		q.Offset = 4 + _offset
	} else {
		q.Offset = _offset + uint32(1+len(c1))
	}
	c2 := q.readArea(q.Offset)
	return c1, c2
}

func (q *Pointer) readArea(offset uint32) []byte {
	q.Offset = offset
	mode := q.readData(1)[0]
	if mode == 0x01 || mode == 0x02 {
		return q.readArea(byteToUInt32(q.readData(3)))
	}
	q.Offset = offset
	return q.readString()
}

func (q *Pointer) readString() []byte {
	data := make([]byte, 0)
	for {
		buf := q.readData(1)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return data
}

func (q *Pointer) searchIndex(ip uint64) uint32 {

	q.ItemLen = 8
	q.IndexLen = 11

	header = dat[8:24]
	start = binary.LittleEndian.Uint32(header[8:])
	counts := binary.LittleEndian.Uint32(header[:8])
	end = start + counts*q.IndexLen

	buf := make([]byte, q.IndexLen)

	for {
		mid := start + q.IndexLen*(((end-start)/q.IndexLen)>>1)
		buf = dat[mid : mid+q.IndexLen]
		_ip := binary.LittleEndian.Uint64(buf[:q.ItemLen])

		if end-start == q.IndexLen {
			if ip >= binary.LittleEndian.Uint64(dat[end:end+q.ItemLen]) {
				buf = dat[end : end+q.IndexLen]
			}
			return byteToUInt32(buf[q.ItemLen:])
		}

		if _ip > ip {
			end = mid
		} else if _ip < ip {
			start = mid
		} else if _ip == ip {
			return byteToUInt32(buf[q.ItemLen:])
		}
	}
}

func byteToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}

// OnlineDownload 在线下载
func (q *Pointer) OnlineDownload() (err error) {
	tmpData, err := getOnline()
	if err != nil {
		return errors.New("下载失败")
	}
	if err := ioutil.WriteFile("./ipv6wry.db", tmpData, 0644); err == nil {
		log.Printf("已下载最新 ZX IPv6数据库 %s ", "./ipv6wry.db")
	} else {
		return errors.New("保存失败")
	}
	return nil
}
