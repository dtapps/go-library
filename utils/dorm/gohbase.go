package dorm

import "github.com/tsuna/gohbase"

type ConfigHbaseClient struct {
	Dns string
}

type HbaseClient struct {
	Db     *gohbase.Client    // 驱动
	config *ConfigHbaseClient // 配置
}

func NewHbaseClient(config *ConfigHbaseClient) (*HbaseClient, error) {

	c := &HbaseClient{config: config}

	db := gohbase.NewClient(c.config.Dns)
	if db != nil {
		panic("Hbase New failed")
	}

	c.Db = &db

	return c, nil
}
