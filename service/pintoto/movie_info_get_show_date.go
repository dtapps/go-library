package pintoto

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetShowDateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		DateList []string `json:"dateList"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetShowDateResult struct {
	Result GetShowDateResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func newGetShowDateResult(result GetShowDateResponse, body []byte, http gorequest.Response, err error) *GetShowDateResult {
	return &GetShowDateResult{Result: result, Body: body, Http: http, Err: err}
}

// GetShowDate 包含某电影的日期 https://www.showdoc.com.cn/1154868044931571/6091788579441818
func (c *Client) GetShowDate(cityId, filmId int) *GetShowDateResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("cityId", cityId)
	param.Set("filmId", filmId)
	// 转换
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(apiUrl+"/movieapi/movie-info/get-show-date", params)
	// 定义
	var response GetShowDateResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGetShowDateResult(response, request.ResponseBody, request, err)
}
