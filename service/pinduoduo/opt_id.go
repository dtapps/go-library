package pinduoduo

type GetOptIdListResponse struct {
	OptId   int64  `json:"opt_id"`
	OptName string `json:"opt_name"`
}

func (c *Client) GetOptIdList() []GetOptIdListResponse {

	var lists []GetOptIdListResponse

	lists = append(lists, GetOptIdListResponse{
		OptId:   1,
		OptName: "食品",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   4,
		OptName: "母婴",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   13,
		OptName: "水果",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   14,
		OptName: "女装",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   15,
		OptName: "百货",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   16,
		OptName: "美妆",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   18,
		OptName: "电器",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   590,
		OptName: "虚拟",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   743,
		OptName: "男装",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   818,
		OptName: "家纺",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   1281,
		OptName: "鞋包",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   1282,
		OptName: "内衣",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   1451,
		OptName: "运动",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   1917,
		OptName: "家装",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   2048,
		OptName: "汽车",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   2478,
		OptName: "文具",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   2974,
		OptName: "家具",
	})
	lists = append(lists, GetOptIdListResponse{
		OptId:   3279,
		OptName: "医药",
	})

	return lists
}
