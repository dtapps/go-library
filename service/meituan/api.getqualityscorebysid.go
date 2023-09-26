package meituan

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"net/http"
)

type ApiGetQuaLitYsCoreBySidResponse struct {
	Status int    `json:"status"`
	Des    string `json:"des"`
	Data   struct {
		DataList []struct {
			Appkey         string `json:"appkey"`         // appkey
			Sid            string `json:"sid"`            // 推广位sid
			Date           string `json:"date"`           // 质量分归属日期
			QualityGrade   string `json:"qualityGrade"`   // 质量分
			RepurchaseRate string `json:"repurchaseRate"` // sid维度的七日复购率
		} `json:"dataList"`
		Total int `json:"total"`
	} `json:"data"`
}
type ApiGetQuaLitYsCoreBySidResult struct {
	Result ApiGetQuaLitYsCoreBySidResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newApiGetQuaLitYsCoreBySidResult(result ApiGetQuaLitYsCoreBySidResponse, body []byte, http gorequest.Response) *ApiGetQuaLitYsCoreBySidResult {
	return &ApiGetQuaLitYsCoreBySidResult{Result: result, Body: body, Http: http}
}

// ApiGetQuaLitYsCoreBySid 优选sid质量分&复购率查询
// https://union.meituan.com/v2/apiDetail?id=28
func (c *Client) ApiGetQuaLitYsCoreBySid(ctx context.Context, notMustParams ...gorequest.Params) (*ApiGetQuaLitYsCoreBySidResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params.Set("ts", gotime.Current().Timestamp())
	params.Set("appkey", c.GetAppKey())
	params.Set("sign", c.getSign(c.GetSecret(), params))
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/getqualityscorebysid", params, http.MethodGet)
	if err != nil {
		return newApiGetQuaLitYsCoreBySidResult(ApiGetQuaLitYsCoreBySidResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiGetQuaLitYsCoreBySidResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiGetQuaLitYsCoreBySidResult(response, request.ResponseBody, request), err
}
