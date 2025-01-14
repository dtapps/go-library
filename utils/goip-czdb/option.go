package goip_czdb

type Option func(*Client)

func WithV4Config(dbFile string, queryType string, key string) Option {
	return func(c *Client) {
		c.option.v4.dbFile = dbFile
		c.option.v4.queryType = queryType
		c.option.v4.key = key
	}
}

func WithV6Config(dbFile string, queryType string, key string) Option {
	return func(c *Client) {
		c.option.v6.dbFile = dbFile
		c.option.v6.queryType = queryType
		c.option.v6.key = key
	}
}
