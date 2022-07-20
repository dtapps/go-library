package wechatopen

import (
	"encoding/json"
	"fmt"
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
	Err    error                  // 错误
}

func newWxaGetCategoryResult(result WxaGetCategoryResponse, body []byte, http gorequest.Response, err error) *WxaGetCategoryResult {
	return &WxaGetCategoryResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetCategory 获取审核时可填写的类目信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/category/get_category.html
func (c *Client) WxaGetCategory() *WxaGetCategoryResult {
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/get_category?access_token=%s", c.GetAuthorizerAccessToken()), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response WxaGetCategoryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaGetCategoryResult(response, request.ResponseBody, request, err)
}
