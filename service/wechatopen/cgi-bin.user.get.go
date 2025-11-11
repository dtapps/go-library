package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type UserGet struct {
	Total int `json:"total"` // 关注该公众账号的总用户数
	Count int `json:"count"` // 拉取的OPENID个数，最大值为10000
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"` // 列表数据，OPENID的列表
	NextOpenid string `json:"next_openid"` // 拉取列表的最后一个用户的OPENID
}

// UserGet 获取用户列表
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Getting_a_User_List.html
func (c *Client) UserGet(ctx context.Context, notMustParams ...*gorequest.Params) (response UserGet, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	err = c.request(ctx, c.WithUrlAuthorizerAccessToken("cgi-bin/user/get"), params, http.MethodPost, &response)
	return
}
