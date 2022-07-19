package dorm

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type ConfigNeo4j4Client struct {
	Dns      string
	Username string
	Password string
	realm    string
}

type Neo4j4Client struct {
	Db     *neo4j.Driver       // 驱动
	config *ConfigNeo4j4Client // 配置
}

func NewNo44Client(config *ConfigNeo4j4Client) (*Neo4j4Client, error) {

	c := &Neo4j4Client{config: config}

	driver, err := neo4j.NewDriver(c.config.Dns, neo4j.BasicAuth(c.config.Username, c.config.Password, c.config.realm))
	if err != nil {
		panic(err)
	}
	defer driver.Close()

	c.Db = &driver

	return c, nil
}
