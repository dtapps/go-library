package kashangwl

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ApiProductCacheResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		GoodsId      uint    `json:"goods_id"`
		ApiGoodsId   int64   `json:"api_goods_id"`
		GoodsName    string  `json:"goods_name"`
		GoodsPrice   float64 `json:"goods_price"`
		PurchaseTips string  `json:"purchase_tips"`
	} `json:"data"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type ApiProductCacheResult struct {
	Result ApiProductCacheResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newApiProductCacheResult(result ApiProductCacheResponse, body []byte, http gorequest.Response) *ApiProductCacheResult {
	return &ApiProductCacheResult{Result: result, Body: body, Http: http}
}

// ApiProductCache [缓存，需托管授权]获取单个商品信息
func (c *Client) ApiProductCache(ctx context.Context, notMustParams ...gorequest.Params) (*ApiProductCacheResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("customer_id", c.GetCustomerId())
	// 请求
	request, err := c.requestCache(ctx, fmt.Sprintf("%s/goods_info", apiUrlCache), params, http.MethodGet)
	if err != nil {
		return newApiProductCacheResult(ApiProductCacheResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiProductCacheResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiProductCacheResult(response, request.ResponseBody, request), err
}
