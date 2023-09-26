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
		Alert struct {
			Status  string `json:"status"`
			Content []struct {
				Pubtimestamp  int       `json:"pubtimestamp"` // 发布时间，单位是 Unix 时间戳
				AlertID       string    `json:"alertId"`      // 预警 ID
				Status        string    `json:"status"`       // 预警信息的状态
				Adcode        string    `json:"adcode"`       // 区域代码
				Location      string    `json:"location"`     // 位置
				Province      string    `json:"province"`     // 省
				City          string    `json:"city"`         // 市
				County        string    `json:"county"`       // 县
				Code          string    `json:"code"`         // 预警代码
				Source        string    `json:"source"`       // 发布单位
				Title         string    `json:"title"`        // 标题
				Description   string    `json:"description"`  // 描述
				RegionID      string    `json:"regionId"`
				Latlon        []float64 `json:"latlon"`
				RequestStatus string    `json:"request_status"`
			} `json:"content"`
			Adcodes []struct {
				Adcode int    `json:"adcode"`
				Name   string `json:"name"`
			} `json:"adcodes"` // 行政区划层级信息
		} `json:"alert"` // 预警数据
		Minutely struct {
			Status          string    `json:"status"`
			Datasource      string    `json:"datasource"`
			Precipitation2H []float64 `json:"precipitation_2h"`
			Precipitation   []float64 `json:"precipitation"`
			Probability     []float64 `json:"probability"`
			Description     string    `json:"description"`
		} `json:"minutely"` // 分钟级预报
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
func (c *Client) Minutely(ctx context.Context, location string, notMustParams ...gorequest.Params) (*MinutelyResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.getApiUrl()+fmt.Sprintf("/%s/minutely", location), params, http.MethodGet)
	if err != nil {
		return newMinutelyResult(MinutelyResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MinutelyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMinutelyResult(response, request.ResponseBody, request), err
}
