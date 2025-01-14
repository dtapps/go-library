package goip_czdb

import (
	"github.com/zhengjianyang/goCzdb"
	"go.dtapp.net/library/utils/gostring"
)

type Client struct {
	v4Client *goCzdb.DbSearcher
	v6Client *goCzdb.DbSearcher
	option   struct {
		v4 struct {
			dbFile    string
			queryType string
			key       string
		}
		v6 struct {
			dbFile    string
			queryType string
			key       string
		}
	}
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	var err error
	if c.option.v4.dbFile != "" && c.option.v4.queryType != "" && c.option.v4.key != "" {
		c.v4Client, err = goCzdb.NewDbSearcher(c.option.v4.dbFile, c.option.v4.queryType, c.option.v4.key)
		if err != nil {
			return nil, err
		}
	}
	if c.option.v6.dbFile != "" && c.option.v6.queryType != "" && c.option.v6.key != "" {
		c.v6Client, err = goCzdb.NewDbSearcher(c.option.v6.dbFile, c.option.v6.queryType, c.option.v6.key)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) Close() {
	if c.v4Client != nil {
		c.v4Client.Close()
	}
	if c.v6Client != nil {
		c.v6Client.Close()
	}
}

type AnalyseResult struct {
	Ip       string `json:"ip"`       // IP
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
	Isp      string `json:"isp"`      // 运营商
}

func (c *Client) Analyse(ip string) (resp AnalyseResult) {
	var search string

	if c.v4Client != nil {
		search, _ = c.v4Client.Search(ip)
	}

	if search == "" && c.v6Client != nil {
		search, _ = c.v6Client.Search(ip)
	}

	resp.Ip = gostring.SpaceAndLineBreak(ip)
	if search != "" {
		split1 := gostring.Split(search, "\t")
		if len(split1) == 2 {
			resp.Isp = gostring.SpaceAndLineBreak(split1[1])
			split2 := gostring.Split(split1[0], "–")
			if len(split2) == 3 {
				resp.Country = gostring.SpaceAndLineBreak(split2[0])
				resp.Province = gostring.SpaceAndLineBreak(split2[1])
				resp.City = gostring.SpaceAndLineBreak(split2[2])
			}
			if len(split2) == 2 {
				resp.Country = gostring.SpaceAndLineBreak(split2[0])
				resp.Province = gostring.SpaceAndLineBreak(split2[1])
			}
			if len(split2) == 1 {
				resp.Country = gostring.SpaceAndLineBreak(split2[0])
			}
		}
	}

	return resp
}
