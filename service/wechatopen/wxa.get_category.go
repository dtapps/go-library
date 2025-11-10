package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaGetCategoryResponse struct {
	APIResponse  // 错误
	CategoryList []struct {
		FirstClass  string `json:"first_class"`  // 一级类目名称
		SecondClass string `json:"second_class"` // 二级类目名称
		ThirdClass  string `json:"third_class"`  // 三级类目名称
		FirstId     int    `json:"first_id"`     // 一级类目的 ID 编号
		SecondId    int    `json:"second_id"`    // 二级类目的 ID 编号
		ThirdId     int    `json:"third_id"`     // 三级类目的 ID 编号
	} `json:"category_list"`
}

// WxaGetCategory 获取审核时可填写的类目信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html
func (c *Client) WxaGetCategory(ctx context.Context, notMustParams ...*gorequest.Params) (response WxaGetCategoryResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/get_category?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodGet, &response)
	return
}
