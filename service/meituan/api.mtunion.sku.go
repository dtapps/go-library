package meituan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"net/http"
)

type ApiMtUnionSkuResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DataList []struct {
			SkuId        string  `json:"skuId"`        // sku编号
			SkuName      string  `json:"skuName"`      // sku名称
			Price        string  `json:"price"`        // 展示价格，单位分
			Pic          float64 `json:"pic"`          // 商品主图
			CategoryId   float64 `json:"categoryId"`   // 商品类目ID
			CategoryName string  `json:"categoryName"` // 商品类目名称
			SalesVolume  float64 `json:"salesVolume"`  // 当前sku销量
		} `json:"dataList"`
		Total int64 `json:"total"` // 商品总数
	} `json:"data"`
}
type ApiMtUnionSkuResult struct {
	Result ApiMtUnionSkuResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func newApiMtUnionSkuResult(result ApiMtUnionSkuResponse, body []byte, http gorequest.Response, err error) *ApiMtUnionSkuResult {
	return &ApiMtUnionSkuResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiMtUnionSku 商品列表查询（新版）
// https://union.meituan.com/v2/apiDetail?id=31
func (c *Client) ApiMtUnionSku(ctx context.Context, notMustParams ...*gorequest.Params) *ApiMtUnionSkuResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params.Set("ts", gotime.Current().Timestamp())
	params.Set("appkey", c.GetAppKey())
	params.Set("sign", c.getSign(c.GetSecret(), params))
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/getqualityscorebysid", params, http.MethodGet)
	// 定义
	var response ApiMtUnionSkuResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiMtUnionSkuResult(response, request.ResponseBody, request, err)
}
