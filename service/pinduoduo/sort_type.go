package pinduoduo

type GetSortTypeListResponse struct {
	SortType int64  `json:"sort_type"`
	SortName string `json:"sort_name"`
}

func (c *Client) GetSortTypeList() []GetSortTypeListResponse {

	var lists []GetSortTypeListResponse

	lists = append(lists, GetSortTypeListResponse{
		SortType: 1,
		SortName: "实时热销榜",
	})
	lists = append(lists, GetSortTypeListResponse{
		SortType: 2,
		SortName: "实时收益榜",
	})

	return lists
}
