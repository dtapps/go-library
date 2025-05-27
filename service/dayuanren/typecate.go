package dayuanren

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type Typecate struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   []struct {
		Id       int64  `json:"id"`        // 产品类型id
		TypeName string `json:"type_name"` // 产品类型名称
		Cate     []struct {
			Id   int64  `json:"id"`        // 分类ID
			Cate string `json:"type_name"` // 分类名称
			Type int64  `json:"type"`      // 产品类型ID
		} `json:"cate"` // 分类列表
	} `json:"data,omitempty"`
}

// Typecate 获取产品类型和产品分类
// https://www.showdoc.com.cn/dyr/9227005390454727
// https://www.kancloud.cn/boyanyun/boyanyun_huafei/3097252
func (c *Client) Typecate(ctx context.Context, notMustParams ...*gorequest.Params) (response Typecate, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID()) // 商户ID

	// 请求
	err = c.request(ctx, "index/typecate", params, &response)
	return
}
