package wikeyun

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error                          // 错误
}

func NewRestRechargeMobileInfoResult(result RestRechargeMobileInfoResponse, body []byte, http gorequest.Response, err error) *RestRechargeMobileInfoResult {
	return &RestRechargeMobileInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargeMobileInfo 查询手机归属地信息以及是否携号转网
// https://open.wikeyun.cn/#/apiDocument/9/document/374
func (c *Client) RestRechargeMobileInfo(orderNumber string) *RestRechargeMobileInfoResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("order_number", orderNumber) // 平台单号
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/rest/Recharge/mobileInfo", params)
	// 定义
	var response RestRechargeMobileInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewRestRechargeMobileInfoResult(response, request.ResponseBody, request, err)
}
