package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetVersionResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

type GetVersionResult struct {
	Result GetVersionResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newGetVersionResult(result GetVersionResponse, body []byte, http gorequest.Response) *GetVersionResult {
	return &GetVersionResult{Result: result, Body: body, Http: http}
}

// GetVersion 获取同步版本号 https://www.showdoc.com.cn/1154868044931571/6566701084841699
func (c *Client) GetVersion(ctx context.Context, notMustParams ...*gorequest.Params) (*GetVersionResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-version", params)
	if err != nil {
		return newGetVersionResult(GetVersionResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetVersionResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetVersionResult(response, request.ResponseBody, request), err
}
