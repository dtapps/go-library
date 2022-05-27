package jd

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type UnionOpenActivityQueryResultResponse struct {
	JdUnionOpenActivityQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_activity_query_responce"`
}

type UnionOpenActivityQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		ActStatus    int    `json:"actStatus"`
		Advantage    string `json:"advantage"`
		CategoryList []struct {
			CategoryId int `json:"categoryId"`
			Type       int `json:"type"`
		} `json:"categoryList,omitempty"`
		Content      string `json:"content"`
		DownloadCode string `json:"downloadCode"`
		DownloadUrl  string `json:"downloadUrl"`
		EndTime      int64  `json:"endTime"`
		Id           int    `json:"id"`
		ImgList      []struct {
			ImgName     string `json:"imgName"`
			ImgUrl      string `json:"imgUrl"`
			WidthHeight string `json:"widthHeight"`
		} `json:"imgList,omitempty"`
		PlatformType       int    `json:"platformType"`
		PromotionEndTime   int64  `json:"promotionEndTime,omitempty"`
		PromotionStartTime int64  `json:"promotionStartTime,omitempty"`
		Recommend          int    `json:"recommend,omitempty"`
		StartTime          int64  `json:"startTime"`
		Tag                string `json:"tag"`
		Title              string `json:"title"`
		UpdateTime         int64  `json:"updateTime"`
		UrlM               string `json:"urlM"`
		UrlPC              string `json:"urlPC"`
	} `json:"data"`
	Message    string `json:"message"`
	RequestId  string `json:"requestId"`
	TotalCount int    `json:"totalCount"`
}

type UnionOpenActivityQueryResult struct {
	Responce UnionOpenActivityQueryResultResponse // 结果
	Result   UnionOpenActivityQueryQueryResult    // 结果
	Body     []byte                               // 内容
	Http     gorequest.Response                   // 请求
	Err      error                                // 错误
}

func NewUnionOpenActivityQueryResult(responce UnionOpenActivityQueryResultResponse, result UnionOpenActivityQueryQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenActivityQueryResult {
	return &UnionOpenActivityQueryResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenActivityQuery 活动查询接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.activity.query
func (app *App) UnionOpenActivityQuery(notMustParams ...Params) *UnionOpenActivityQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.activity.query", notMustParams...)
	// 请求
	request, err := app.request(params)
	// 定义
	var responce UnionOpenActivityQueryResultResponse
	var result UnionOpenActivityQueryQueryResult
	err = json.Unmarshal(request.ResponseBody, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenActivityQueryResponce.QueryResult), &result)
	return NewUnionOpenActivityQueryResult(responce, result, request.ResponseBody, request, err)
}
