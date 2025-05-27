package dayuanren

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type User struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   struct {
		Id       int64  `json:"id"`       // userid
		Username string `json:"username"` // 名称
		Balance  string `json:"balance"`  // 余额
	} `json:"data,omitempty"`
}

// User 查询用户信息
// https://www.showdoc.com.cn/dyr/9227004018562421
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3097251
func (c *Client) User(ctx context.Context, notMustParams ...*gorequest.Params) (response User, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID()) // 账号ID

	// 请求
	err = c.request(ctx, "index/user", params, &response)
	return
}
