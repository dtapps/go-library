package pinduoduo

type GetResourceTypeListResponse struct {
	ResourceType int64  `json:"resource_type"`
	ResourceName string `json:"resource_name"`
}

func (c *Client) GetResourceTypeList() []GetResourceTypeListResponse {

	var lists []GetResourceTypeListResponse

	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 4,
		ResourceName: "限时秒杀",
	})
	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 39997,
		ResourceName: "充值中心",
	})
	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 39998,
		ResourceName: "活动转链",
	})
	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 39996,
		ResourceName: "百亿补贴",
	})
	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 39999,
		ResourceName: "电器城",
	})
	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 40000,
		ResourceName: "领券中心",
	})
	lists = append(lists, GetResourceTypeListResponse{
		ResourceType: 50005,
		ResourceName: "火车票",
	})

	return lists
}
