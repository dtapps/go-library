package pinduoduo

type GetCatIdListResponse struct {
	CatId   int64  `json:"cat_id"`
	CatName string `json:"cat_name"`
}

func (c *Client) GetCatIdList() []GetCatIdListResponse {

	var lists []GetCatIdListResponse

	lists = append(lists, GetCatIdListResponse{
		CatId:   20100,
		CatName: "百货",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20200,
		CatName: "母婴",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20300,
		CatName: "食品",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20400,
		CatName: "女装",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20500,
		CatName: "电器",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20600,
		CatName: "鞋包",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20700,
		CatName: "内衣",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20800,
		CatName: "美妆",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   20900,
		CatName: "男装",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21000,
		CatName: "水果",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21100,
		CatName: "家纺",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21200,
		CatName: "文具",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21300,
		CatName: "运动",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21400,
		CatName: "虚拟",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21500,
		CatName: "汽车",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21600,
		CatName: "家装",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21700,
		CatName: "家具",
	})
	lists = append(lists, GetCatIdListResponse{
		CatId:   21800,
		CatName: "医药",
	})

	return lists
}
