package czdb

import (
	"fmt"

	"github.com/tagphi/czdb-search-golang/pkg/db"
)

// Client 客户端
type Client struct {
	v4Filepath string
	v4Db       *db.DBSearcher
	v6Filepath string
	v6Db       *db.DBSearcher
}

// New 创建
func New(v4Filepath string, v6Filepath string, key string, searchType db.SearchType) (*Client, error) {

	var err error
	c := &Client{
		v4Filepath: v4Filepath,
		v6Filepath: v6Filepath,
	}

	if c.v4Filepath == "" && c.v6Filepath == "" {
		return nil, fmt.Errorf("v4Filepath and v6Filepath are empty")
	}

	if v4Filepath != "" {
		c.v4Db, err = db.InitDBSearcher(v4Filepath, key, searchType)
		if err != nil {
			return nil, err
		}
	}

	if v6Filepath != "" {
		c.v6Db, err = db.InitDBSearcher(v6Filepath, key, searchType)
		if err != nil {
			return nil, err
		}
	}

	return c, err
}

// Close 关闭
func (c *Client) Close() {
	if c.v4Filepath != "" {
		db.CloseDBSearcher(c.v4Db)
	}
	if c.v6Filepath != "" {
		db.CloseDBSearcher(c.v6Db)
	}
}
