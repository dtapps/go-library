package eastiot

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IotApiQueryOrderedPkgInfoResponse struct {
	Code   int64 `json:"code"`
	Istest int64 `json:"istest"`
	Data   []struct {
		Name      string  `json:"name"`      // 流量包名字
		PkgId     int64   `json:"pkgId"`     // 流量包ID
		Traffic   int64   `json:"traffic"`   // 流量大小，单位:MB
		Ntraffic  float64 `json:"ntraffic"`  // 已用量，单位:MB
		Starttime int64   `json:"starttime"` // 流量生效起始时间时间戳
		Endtime   int64   `json:"endtime"`   // 流量生效结束时间时间戳
		Addtime   int64   `json:"addtime"`   // 订购时间时间戳
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiQueryOrderedPkgInfoResult struct {
	Result IotApiQueryOrderedPkgInfoResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newIotApiQueryOrderedPkgInfoResult(result IotApiQueryOrderedPkgInfoResponse, body []byte, http gorequest.Response) *IotApiQueryOrderedPkgInfoResult {
	return &IotApiQueryOrderedPkgInfoResult{Result: result, Body: body, Http: http}
}

// IotApiQueryOrderedPkgInfo 查询流量卡已订购流量包
// https://www.showdoc.com.cn/916774523755909/5092045889939625
func (c *Client) IotApiQueryOrderedPkgInfo(ctx context.Context, simId string, notMustParams ...gorequest.Params) (*IotApiQueryOrderedPkgInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("simId", simId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/Api/IotApi/queryOrderedPkgInfo", params, http.MethodPost)
	if err != nil {
		return newIotApiQueryOrderedPkgInfoResult(IotApiQueryOrderedPkgInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response IotApiQueryOrderedPkgInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIotApiQueryOrderedPkgInfoResult(response, request.ResponseBody, request), err
}
