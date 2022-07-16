package dorm

import (
	_ "github.com/go-sql-driver/mysql"
)

func NewEntMysqlClient(config *ConfigEntClient) (*EntClient, error) {

	//var err error
	//c := &EntClient{config: config}

	//client, err := ent.Open("mysql", c.config.Dns)
	//if err != nil {
	//	return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	//}
	//defer client.Close()

	return nil, nil
}
