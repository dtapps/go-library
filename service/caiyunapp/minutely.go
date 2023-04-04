package caiyunapp

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type MinutelyResponse struct {
	Status     string    `json:"status"`
	ApiVersion string    `json:"api_version"`
	ApiStatus  string    `json:"api_status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	Tzshift    float64   `json:"tzshift"`
	Timezone   string    `json:"timezone"`
	ServerTime float64   `json:"server_time"`
	Location   []float64 `json:"location"`
	Result     struct {
		Minutely struct {
			Status          string    `json:"status"`
			Datasource      string    `json:"datasource"`
			Precipitation2H []float64 `json:"precipitation_2h"`
			Precipitation   []float64 `json:"precipitation"`
			Probability     []float64 `json:"probability"`
			Description     string    `json:"description"`
		} `json:"minutely"`
		Primary          float64 `json:"primary"`
		ForecastKeypoint string  `json:"forecast_keypoint"`
	} `json:"result"`
}

type MinutelyResult struct {
	Result MinutelyResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newMinutelyResult(result MinutelyResponse, body []byte, http gorequest.Response) *MinutelyResult {
	return &MinutelyResult{Result: result, Body: body, Http: http}
}

// Minutely 分钟级预报
// https://docs.caiyunapp.com/docs/minutely
func (c *Client) Minutely(ctx context.Context, locationLongitude, locationLatitude string, notMustParams ...gorequest.Params) (*MinutelyResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s,%s/minutely", locationLatitude, locationLongitude), params, http.MethodGet)
	if err != nil {
		return newMinutelyResult(MinutelyResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MinutelyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMinutelyResult(response, request.ResponseBody, request), err
}
