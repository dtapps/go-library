package pintoto

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
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
	Err    error              // 错误
}

func NewGetVersionResult(result GetVersionResponse, body []byte, http gorequest.Response, err error) *GetVersionResult {
	return &GetVersionResult{Result: result, Body: body, Http: http, Err: err}
}

// GetVersion 获取同步版本号 https://www.showdoc.com.cn/1154868044931571/6566701084841699
func (app *App) GetVersion() *GetVersionResult {
	request, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-version", map[string]interface{}{})
	// 定义
	var response GetVersionResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewGetVersionResult(response, request.ResponseBody, request, err)
}
