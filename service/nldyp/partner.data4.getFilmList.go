package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetFilmListResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Id           string      `json:"id"`           // 影片 id
		FilmCode     string      `json:"filmCode"`     // 影片编码（标准影片 编码去除第 4 位）
		FilmName     string      `json:"filmName"`     // 影片名称
		Version      string      `json:"version"`      // 影片制式
		Duration     int         `json:"duration"`     // 影片时长单位分钟
		PublishDate  int         `json:"publishDate"`  // 上映时间 unix 时间 戳
		Director     string      `json:"director"`     // 导演
		CastType     int         `json:"castType"`     // 0、主演 1、配音演 员
		Cast         string      `json:"cast"`         // 主演/配音演员
		Introduction string      `json:"introduction"` // 影片描述
		Wantview     int         `json:"wantview"`     // 想看人数
		Score        int         `json:"score"`        // 电影评分
		Cover        string      `json:"cover"`        // 封面图
		Area         string      `json:"area"`         // 影片归属地区
		Type         string      `json:"type"`         // 影片类型
		PlanNum      interface{} `json:"planNum"`
		PreSaleFlag  int         `json:"preSaleFlag"` // 是 否 预 售 0 、 否 1、预售
	} `json:"data"`
}

type PartnerData4GetFilmListResult struct {
	Result PartnerData4GetFilmListResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newPartnerData4GetFilmListResult(result PartnerData4GetFilmListResponse, body []byte, http gorequest.Response) *PartnerData4GetFilmListResult {
	return &PartnerData4GetFilmListResult{Result: result, Body: body, Http: http}
}

// PartnerData4GetFilmList 获取影片
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=b13d7051-6a31-49d4-ba49-42e423da41d3
func (c *Client) PartnerData4GetFilmList(ctx context.Context, notMustParams ...gorequest.Params) (*PartnerData4GetFilmListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getFilmList", params)
	if err != nil {
		return newPartnerData4GetFilmListResult(PartnerData4GetFilmListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4GetFilmListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetFilmListResult(response, request.ResponseBody, request), err
}
