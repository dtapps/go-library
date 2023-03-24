package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGetCategoryResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	CategoryList []struct {
		FirstClass  string `json:"first_class"`  // 一级类目名称
		SecondClass string `json:"second_class"` // 二级类目名称
		ThirdClass  string `json:"third_class"`  // 三级类目名称
		FirstId     int    `json:"first_id"`     // 一级类目的 ID 编号
		SecondId    int    `json:"second_id"`    // 二级类目的 ID 编号
		ThirdId     int    `json:"third_id"`     // 三级类目的 ID 编号
	} `json:"category_list"`
}

type WxaGetCategoryResult struct {
	Result WxaGetCategoryResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newWxaGetCategoryResult(result WxaGetCategoryResponse, body []byte, http gorequest.Response) *WxaGetCategoryResult {
	return &WxaGetCategoryResult{Result: result, Body: body, Http: http}
}

// WxaGetCategory 获取审核时可填写的类目信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html
func (c *Client) WxaGetCategory(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGetCategoryResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/get_category?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodGet)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaGetCategoryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaGetCategoryResult(response, request.ResponseBody, request), nil
}
