package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetCinemaResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Id               string      `json:"id"`           // 影城 id （注：用于请求其他数据）
		CinemaId         int         `json:"cinemaId"`     // 版本4 影城ID （注：用于锁座，判断是否有秒出票，为空则不能秒出票）
		CinemaCode       int         `json:"cinemaCode"`   // 影城编码
		CinemaName       string      `json:"cinemaName"`   // 影城名称
		ProvinceId       string      `json:"provinceId"`   // 省份ID
		CityId           string      `json:"cityId"`       // 城市ID
		CountyId         int         `json:"countyId"`     // 区县ID
		Address          string      `json:"address"`      // 影城地址
		Longitude        string      `json:"longitude"`    // 经度(高德)
		Latitude         string      `json:"latitude"`     // 维度(高德)
		Province         string      `json:"province"`     // 省份
		City             string      `json:"city"`         // 城市
		County           string      `json:"county"`       // 区县
		StopSaleTime     string      `json:"stopSaleTime"` // 停售时间
		Direct           interface{} `json:"direct"`
		BackTicketConfig interface{} `json:"backTicketConfig"`
	} `json:"data"`
}

type PartnerData4GetCinemaResult struct {
	Result PartnerData4GetCinemaResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
}

func newPartnerData4GetCinemaResult(result PartnerData4GetCinemaResponse, body []byte, http gorequest.Response) *PartnerData4GetCinemaResult {
	return &PartnerData4GetCinemaResult{Result: result, Body: body, Http: http}
}

// PartnerData4GetCinema 获取影院
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=a1db1066-f346-4f9a-bffe-b892b3c73a1d
func (c *Client) PartnerData4GetCinema(ctx context.Context, notMustParams ...gorequest.Params) (*PartnerData4GetCinemaResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getCinema", params)
	if err != nil {
		return newPartnerData4GetCinemaResult(PartnerData4GetCinemaResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4GetCinemaResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetCinemaResult(response, request.ResponseBody, request), err
}
