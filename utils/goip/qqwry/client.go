package qqwry

import (
	_ "embed"
	"encoding/binary"
	"log"
)

var (
	header  []byte
	country []byte
	area    []byte
	offset  uint32
	start   uint32
	end     uint32
)

//go:embed qqwry.dat
var datBuff []byte

type Client struct {
	Offset   uint32
	ItemLen  uint32
	IndexLen uint32
}

func New() *Client {

	c := &Client{}

	buf := datBuff[0:8]
	start := binary.LittleEndian.Uint32(buf[:4])
	end := binary.LittleEndian.Uint32(buf[4:])

	num := int64((end-start)/7 + 1)
	log.Printf("qqwry.dat 共加载：%d 条ip记录\n", num)

	return c
}

// ReadData 从文件中读取数据
func (c *Client) readData(length uint32) (rs []byte) {
	end := c.Offset + length
	dataNum := uint32(len(datBuff))
	if c.Offset > dataNum {
		return nil
	}

	if end > dataNum {
		end = dataNum
	}
	rs = datBuff[c.Offset:end]
	c.Offset = end
	return rs
}

// 获取地址信息
func (c *Client) getAddr() ([]byte, []byte) {
	mode := c.readData(1)[0]
	if mode == 0x01 {
		// [IP][0x01][国家和地区信息的绝对偏移地址]
		c.Offset = byteToUInt32(c.readData(3))
		return c.getAddr()
	}
	// [IP][0x02][信息的绝对偏移][...] or [IP][国家][...]
	_offset := c.Offset - 1
	c1 := c.readArea(_offset)
	if mode == 0x02 {
		c.Offset = 4 + _offset
	} else {
		c.Offset = _offset + uint32(1+len(c1))
	}
	c2 := c.readArea(c.Offset)
	return c1, c2
}

// 读取区
func (c *Client) readArea(offset uint32) []byte {
	c.Offset = offset
	mode := c.readData(1)[0]
	if mode == 0x01 || mode == 0x02 {
		return c.readArea(byteToUInt32(c.readData(3)))
	}
	c.Offset = offset
	return c.readString()
}

// 读取字符串
func (c *Client) readString() []byte {
	data := make([]byte, 0)
	for {
		buf := c.readData(1)
		if buf[0] == 0 {
			break
		}
		data = append(data, buf[0])
	}
	return data
}

// 搜索索引
func (c *Client) searchIndex(ip uint32) uint32 {
	c.ItemLen = 4
	c.IndexLen = 7
	header = datBuff[0:8]
	start = binary.LittleEndian.Uint32(header[:4])
	end = binary.LittleEndian.Uint32(header[4:])

	buf := make([]byte, c.IndexLen)

	for {
		mid := start + c.IndexLen*(((end-start)/c.IndexLen)>>1)
		buf = datBuff[mid : mid+c.IndexLen]
		_ip := binary.LittleEndian.Uint32(buf[:c.ItemLen])

		if end-start == c.IndexLen {
			if ip >= binary.LittleEndian.Uint32(datBuff[end:end+c.ItemLen]) {
				buf = datBuff[end : end+c.IndexLen]
			}
			return byteToUInt32(buf[c.ItemLen:])
		}

		if _ip > ip {
			end = mid
		} else if _ip < ip {
			start = mid
		} else if _ip == ip {
			return byteToUInt32(buf[c.ItemLen:])
		}
	}
}

// 字节转UInt32
func byteToUInt32(data []byte) uint32 {
	i := uint32(data[0]) & 0xff
	i |= (uint32(data[1]) << 8) & 0xff00
	i |= (uint32(data[2]) << 16) & 0xff0000
	return i
}
