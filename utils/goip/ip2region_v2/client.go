package ip2region_v2

import _ "embed"

//go:embed ip2region.xdb
var cBuff []byte

type Client struct {
	db *Searcher
}

func New() (*Client, error) {

	var err error
	c := &Client{}

	// 1、从 dbPath 加载整个 xdb 到内存

	// 2、用全局的 cBuff 创建完全基于内存的查询对象。
	c.db, err = NewWithBuffer(cBuff)
	if err != nil {
		return nil, err
	}

	return c, err
}
