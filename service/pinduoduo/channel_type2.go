package pinduoduo

type GetChannel2TypeListResponse struct {
	Channel2Type int64  `json:"channel_2_type"`
	Channel2Name string `json:"channel_2_name"`
}

func (c *Client) GetChannel2TypeList() []GetChannel2TypeListResponse {

	var lists []GetChannel2TypeListResponse

	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 0,
		Channel2Name: "红包",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 1,
		Channel2Name: "活动列表",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 2,
		Channel2Name: "新人红包",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 3,
		Channel2Name: "刮刮卡",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 5,
		Channel2Name: "员工内购",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 10,
		Channel2Name: "生成绑定备案链接",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 12,
		Channel2Name: "砸金蛋",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 14,
		Channel2Name: "千万补贴B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 15,
		Channel2Name: "充值中心B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 16,
		Channel2Name: "千万补贴C端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 17,
		Channel2Name: "千万补贴投票页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 23,
		Channel2Name: "超级红包",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 24,
		Channel2Name: "礼金全场N折活动B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 27,
		Channel2Name: "带货赢千万",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 28,
		Channel2Name: "满减券活动B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 29,
		Channel2Name: "满减券活动C端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 30,
		Channel2Name: "免单B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 31,
		Channel2Name: "免单C端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 32,
		Channel2Name: "转盘得现金B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 33,
		Channel2Name: "转盘得现金C端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 34,
		Channel2Name: "千万神券C端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 35,
		Channel2Name: "千万神券B端页面",
	})
	lists = append(lists, GetChannel2TypeListResponse{
		Channel2Type: 37,
		Channel2Name: "超级红包B端推品页",
	})

	return lists
}
