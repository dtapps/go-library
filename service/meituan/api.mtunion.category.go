package meituan

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"net/http"
)

type ApiMtUnionCategoryResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DataList []struct {
			CategoryId   float64 `json:"categoryId"`   // 商品类目ID
			CategoryName string  `json:"categoryName"` // 商品类目名称
		} `json:"dataList"`
		Total int64 `json:"total"` // 查询总数
	} `json:"data"`
}
type ApiMtUnionCategoryResult struct {
	Result ApiMtUnionCategoryResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func newApiMtUnionCategoryResult(result ApiMtUnionCategoryResponse, body []byte, http gorequest.Response, err error) *ApiMtUnionCategoryResult {
	return &ApiMtUnionCategoryResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiMtUnionCategory 商品类目查询（新版）
// https://union.meituan.com/v2/apiDetail?id=30
func (c *Client) ApiMtUnionCategory(ctx context.Context, notMustParams ...gorequest.Params) *ApiMtUnionCategoryResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["ts"] = gotime.Current().Timestamp()
	params["appkey"] = c.GetAppKey()
	params["sign"] = c.getSign(c.GetSecret(), params)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/getqualityscorebysid", params, http.MethodGet)
	// 定义
	var response ApiMtUnionCategoryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiMtUnionCategoryResult(response, request.ResponseBody, request, err)
}
