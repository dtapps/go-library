package eastiot

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type IotApiGetAllSimTypeResponse struct {
	Code int64 `json:"code"`
	Data []struct {
		Type   int64  `json:"type"`   // 卡类型
		Name   string `json:"name"`   // 类型名
		MOrder int64  `json:"mOrder"` // 是否支持单次充值多个流量包，0:不支持 1:支持
	} `json:"data"`
	Msg string `json:"msg"`
}

type IotApiGetAllSimTypeResult struct {
	Result IotApiGetAllSimTypeResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newIotApiGetAllSimTypeResult(result IotApiGetAllSimTypeResponse, body []byte, http gorequest.Response) *IotApiGetAllSimTypeResult {
	return &IotApiGetAllSimTypeResult{Result: result, Body: body, Http: http}
}

// IotApiGetAllSimType 卡类型列表查询
// https://www.showdoc.com.cn/916774523755909/4858492092033167
func (c *Client) IotApiGetAllSimType(ctx context.Context, notMustParams ...*gorequest.Params) (*IotApiGetAllSimTypeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/Api/IotApi/getAllSimType", params, http.MethodPost)
	if err != nil {
		return newIotApiGetAllSimTypeResult(IotApiGetAllSimTypeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response IotApiGetAllSimTypeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newIotApiGetAllSimTypeResult(response, request.ResponseBody, request), err
}
