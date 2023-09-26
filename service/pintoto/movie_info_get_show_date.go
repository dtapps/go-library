package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
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
}

func newGetShowDateResult(result GetShowDateResponse, body []byte, http gorequest.Response) *GetShowDateResult {
	return &GetShowDateResult{Result: result, Body: body, Http: http}
}

// GetShowDate 包含某电影的日期 https://www.showdoc.com.cn/1154868044931571/6091788579441818
func (c *Client) GetShowDate(ctx context.Context, cityId, filmId int, notMustParams ...gorequest.Params) (*GetShowDateResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cityId", cityId)
	params.Set("filmId", filmId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-show-date", params)
	if err != nil {
		return newGetShowDateResult(GetShowDateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetShowDateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetShowDateResult(response, request.ResponseBody, request), err
}
