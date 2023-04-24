package eastiot

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IotApiRechargeSimResponse struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
}

type IotApiRechargeSimResult struct {
	Result IotApiRechargeSimResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func newIotApiRechargeSimResult(result IotApiRechargeSimResponse, body []byte, http gorequest.Response, err error) *IotApiRechargeSimResult {
	return &IotApiRechargeSimResult{Result: result, Body: body, Http: http, Err: err}
}

// IotApiRechargeSim 单卡流量充值
// https://www.showdoc.com.cn/916774523755909/4880284631482420
func (c *Client) IotApiRechargeSim(ctx context.Context, notMustParams ...gorequest.Params) *IotApiRechargeSimResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/Api/IotApi/rechargeSim", params, http.MethodPost)
	// 定义
	var response IotApiRechargeSimResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIotApiRechargeSimResult(response, request.ResponseBody, request, err)
}
