package meituan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"net/http"
)

type ApiMtUnionPoiResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		DataList []struct {
			PoiViewId           string  `json:"poiViewId"`           // POI门店ID
			PoiName             string  `json:"poiName"`             // POI名称
			PoiPicUrl           string  `json:"poiPicUrl"`           // 店铺图URL
			PoiScore            string  `json:"poiScore"`            // 店铺评分，满分5分
			MonthSale           string  `json:"monthSale"`           // 月售量
			ShippingFee         string  `json:"shippingFee"`         // 配送费金额，单位元
			MinPrice            string  `json:"minPrice"`            // 起送金额，单位元
			Distance            string  `json:"distance"`            // 门店距离，单位米
			AvgDeliveryTime     string  `json:"avgDeliveryTime"`     // 配送时长，单位分钟
			ReduceShippingFee   float64 `json:"reduceShippingFee"`   // 满减配送费
			PoiMarkTagUrl       string  `json:"poiMarkTagUrl"`       // 角标信息
			MerchantFullSale    string  `json:"merchantFullSale"`    // 店铺满减,举例：38减25
			MerchantDiscount    string  `json:"merchantDiscount"`    // 店铺折扣，举例：3.4折起
			NewCustomerDiscount string  `json:"newCustomerDiscount"` // 新客立减，举例：新客减1
			RebateCoupon        string  `json:"rebateCoupon"`        // 返券，举例：返3元券
			MerchantCoupon      string  `json:"merchantCoupon"`      // 商家券，举例：领3元券
			FullComplimentary   string  `json:"fullComplimentary"`   // 满赠，举例：满68元得赠品
		} `json:"dataList"`
		PageTraceId string `json:"pageTraceId"` // 分页查询参数，第二次查询传回
	} `json:"data"`
}
type ApiMtUnionPoiResult struct {
	Result ApiMtUnionPoiResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newApiMtUnionPoiResult(result ApiMtUnionPoiResponse, body []byte, http gorequest.Response) *ApiMtUnionPoiResult {
	return &ApiMtUnionPoiResult{Result: result, Body: body, Http: http}
}

// ApiMtUnionPoi 门店POI查询（新版）
// https://union.meituan.com/v2/apiDetail?id=32
func (c *Client) ApiMtUnionPoi(ctx context.Context, notMustParams ...*gorequest.Params) (*ApiMtUnionPoiResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params.Set("ts", gotime.Current().Timestamp())
	params.Set("appkey", c.GetAppKey())
	params.Set("sign", c.getSign(c.GetSecret(), params))
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/getqualityscorebysid", params, http.MethodGet)
	if err != nil {
		return newApiMtUnionPoiResult(ApiMtUnionPoiResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiMtUnionPoiResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiMtUnionPoiResult(response, request.ResponseBody, request), err
}
