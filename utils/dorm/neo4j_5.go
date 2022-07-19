package dorm

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ConfigNeo4j5Client struct {
	Dns      string
	Username string
	Password string
	realm    string
}

type Neo4j5Client struct {
	Db     *neo4j.DriverWithContext // 驱动
	config *ConfigNeo4j5Client      // 配置
}

func NewNeo4j5Client(config *ConfigNeo4j5Client) (*Neo4j5Client, error) {

	c := &Neo4j5Client{config: config}

	driver, err := neo4j.NewDriverWithContext(c.config.Dns, neo4j.BasicAuth(c.config.Username, c.config.Password, c.config.realm))
	if err != nil {
		panic(err)
	}
	defer driver.Close(context.Background())

	c.Db = &driver

	return c, nil
}
