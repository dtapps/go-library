package eastiot

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IotApiQueryUserPkgInfoResponse struct {
	Code int64 `json:"code"`
	Data []struct {
		Type    int64   `json:"type"`
		PkgId   int64   `json:"pkgId"`
		PkgName string  `json:"pkgName"`
		Price   float64 `json:"price"`
		Sprice  float64 `json:"sprice"`
		Traffic int64   `json:"traffic"`
		Caltype int64   `json:"caltype"`
		SimType int64   `json:"simType"`
		Isdm    int64   `json:"isdm"`
		Isnm    int64   `json:"isnm"`
		Istest  int64   `json:"istest"`
		Isimm   int64   `json:"isimm"`
		Daynum  int64   `json:"daynum"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQueryUserPkgInfoResult struct {
	Result IotApiQueryUserPkgInfoResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
	Err    error                          // 错误
}

func newIotApiQueryUserPkgInfoResult(result IotApiQueryUserPkgInfoResponse, body []byte, http gorequest.Response, err error) *IotApiQueryUserPkgInfoResult {
	return &IotApiQueryUserPkgInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// IotApiQueryUserPkgInfo 账户可用流量包查询
// https://www.showdoc.com.cn/916774523755909/4850094776758927
func (c *Client) IotApiQueryUserPkgInfo(ctx context.Context) *IotApiQueryUserPkgInfoResult {
	// 请求
	request, err := c.request(ctx, apiUrl+"/Api/IotApi/queryUserPkgInfo", map[string]interface{}{}, http.MethodPost)
	// 定义
	var response IotApiQueryUserPkgInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIotApiQueryUserPkgInfoResult(response, request.ResponseBody, request, err)
}
