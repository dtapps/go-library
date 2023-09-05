package ip2region_v2

import "os"

var cBuff []byte

type Client struct {
	db *Searcher
}

func New(filepath string) (*Client, error) {

	var err error
	c := &Client{}

	// 1、从 dbPath 加载整个 xdb 到内存
	cBuff, err = os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	c.db, err = NewWithBuffer(cBuff)
	if err != nil {
		return nil, err
	}

	return c, err
}

func NewBuff(file []byte) (*Client, error) {

	var err error
	c := &Client{}

	// 1、从 dbPath 加载整个 xdb 到内存
	cBuff = file

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	c.db, err = NewWithBuffer(cBuff)
	if err != nil {
		return nil, err
	}

	return c, err
}
