package pinduoduo

type GetChannelTypeListResponse struct {
	ChannelType int64  `json:"channel_type"`
	ChannelName string `json:"channel_name"`
}

func (c *Client) GetChannelTypeList() []GetChannelTypeListResponse {

	var lists []GetChannelTypeListResponse

	lists = append(lists, GetChannelTypeListResponse{
		ChannelType: 1,
		ChannelName: "今日热销榜",
	})
	lists = append(lists, GetChannelTypeListResponse{
		ChannelType: 3,
		ChannelName: "相似商品推荐",
	})
	lists = append(lists, GetChannelTypeListResponse{
		ChannelType: 4,
		ChannelName: "猜你喜欢",
	})
	lists = append(lists, GetChannelTypeListResponse{
		ChannelType: 5,
		ChannelName: "实时热销榜",
	})
	lists = append(lists, GetChannelTypeListResponse{
		ChannelType: 6,
		ChannelName: "实时收益榜",
	})

	return lists
}
