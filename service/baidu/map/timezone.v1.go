package _map

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type TimezoneV1Response struct {
	Status     int64  `json:"status"`      // 本次API访问状态，如果成功返回0，如果失败返回其他数字。
	TimezoneId string `json:"timezone_id"` // 所在时区ID字符串
	DstOffset  int64  `json:"dst_offset"`  // 夏令时(Daylight Saving Time：DST)时间偏移秒数
	RawOffset  int64  `json:"raw_offset"`  // 坐标点位置时间较协调世界时偏移秒数
}

type TimezoneV1Result struct {
	Result TimezoneV1Response // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newTimezoneV1Result(result TimezoneV1Response, body []byte, http gorequest.Response) *TimezoneV1Result {
	return &TimezoneV1Result{Result: result, Body: body, Http: http}
}

// TimezoneV1 时区服务
// location => 需查询时区的位置坐标 （纬度、经度）,当前仅支持全球陆地坐标查询，海域坐标暂不支持。
// coord_type => 请求参数中坐标的类型，wgs84即GPS经纬度，gcj02即国测局经纬度坐标，bd09ll即百度经纬度坐标，bd09mc即百度米制坐标
// timestamp => 所需时间（用于判断夏令时）。以协调世界时 1970 年 1 月 1 日午夜以来的秒数表示（即Unix时间戳）
// https://lbsyun.baidu.com/faq/api?title=webapi/guide/timezone-base
func (c *Client) TimezoneV1(ctx context.Context, notMustParams ...gorequest.Params) (*TimezoneV1Result, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("ak", c.ak)
	params.Set("output", "json")
	// 请求
	request, err := c.request(ctx, "timezone/v1", params, http.MethodGet)
	if err != nil {
		return newTimezoneV1Result(TimezoneV1Response{}, request.ResponseBody, request), err
	}
	// 定义
	var response TimezoneV1Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTimezoneV1Result(response, request.ResponseBody, request), err
}
