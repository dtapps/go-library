package juhe

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type MobileGetResponse struct {
	Resultcode string `json:"resultcode"` // 返回码
	Reason     string `json:"reason"`     // 返回说明
	Result     struct {
		Province string `json:"province"` // 省份
		City     string `json:"city"`     // 城市，(部分记录可能为空)
		Areacode string `json:"areacode"` // 区号，(部分记录可能为空)
		Zip      string `json:"zip"`      // 邮编，(部分记录可能为空)
		Company  string `json:"company"`  // 运营商
		Card     string `json:"card,omitempty"`
	} `json:"result"` // 返回结果集
}

type MobileGetResult struct {
	Result MobileGetResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newMobileGetResult(result MobileGetResponse, body []byte, http gorequest.Response) *MobileGetResult {
	return &MobileGetResult{Result: result, Body: body, Http: http}
}

// MobileGet 手机号码归属地
// https://www.juhe.cn/docs/api/id/11
func (c *Client) MobileGet(ctx context.Context, phone string, key string, notMustParams ...gorequest.Params) (*MobileGetResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "mobile/get")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("phone", phone)  // 需要查询的手机号码或手机号码前7位
	params.Set("key", key)      // 在个人中心->我的数据,接口名称上方查看
	params.Set("dtype", "json") // 返回数据的格式,xml或json，默认json

	// 请求
	var response MobileGetResponse
	request, err := c.request(ctx, "mobile/get", params, http.MethodGet, &response)
	return newMobileGetResult(response, request.ResponseBody, request), err
}
