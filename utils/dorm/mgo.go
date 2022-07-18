package dorm

import (
	"gopkg.in/mgo.v2"
)

type ConfigMgoClient struct {
	Dns          string
	DatabaseName string // 库名
}

type MgoClient struct {
	config *ConfigMgoClient // 配置
}

func NewMgoClient(config *ConfigMgoClient) (*MgoClient, error) {

	c := &MgoClient{config: config}

	session, err := mgo.Dial(c.config.Dns)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	return c, nil
}
