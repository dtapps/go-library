package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestRechargeMobileInfoResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Status string `json:"status"`
		Oisp   string `json:"oisp"`
		Nisp   string `json:"nisp"`
		Number string `json:"number"`
		Extend struct {
			AreaNum       string `json:"area_num"`
			Isp           string `json:"isp"`
			Prov          string `json:"prov"`
			City          string `json:"city"`
			PostCode      string `json:"post_code"`
			AreaCode      string `json:"area_code"`
			CardProvCode  string `json:"card_prov_code"`
			CardCityCode  string `json:"card_city_code"`
			Lng           string `json:"lng"`
			Lat           string `json:"lat"`
			CityCode      string `json:"city_code"`
			CityShortCode string `json:"city_short_code"`
		} `json:"extend"`
	} `json:"data"`
}

type RestRechargeMobileInfoResult struct {
	Result RestRechargeMobileInfoResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newRestRechargeMobileInfoResult(result RestRechargeMobileInfoResponse, body []byte, http gorequest.Response) *RestRechargeMobileInfoResult {
	return &RestRechargeMobileInfoResult{Result: result, Body: body, Http: http}
}

// RestRechargeMobileInfo 查询手机归属地信息以及是否携号转网
// https://open.wikeyun.cn/#/apiDocument/9/document/374
func (c *Client) RestRechargeMobileInfo(ctx context.Context, orderNumber string, notMustParams ...*gorequest.Params) (*RestRechargeMobileInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("order_number", orderNumber) // 平台单号
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Recharge/mobileInfo", params)
	if err != nil {
		return newRestRechargeMobileInfoResult(RestRechargeMobileInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestRechargeMobileInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestRechargeMobileInfoResult(response, request.ResponseBody, request), err
}
