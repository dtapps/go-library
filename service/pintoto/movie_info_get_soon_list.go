package pintoto

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GetSoonListResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		HasMore int `json:"hasMore"`
		List    []struct {
			Director     string      `json:"director"`     // 导演
			PublishDate  string      `json:"publishDate"`  // 影片上映日期
			VersionTypes string      `json:"versionTypes"` // 上映类型
			Language     string      `json:"language"`     // 语言
			ShowStatus   int         `json:"showStatus"`   // 放映状态：1 正在热映。2 即将上映
			Pic          string      `json:"pic"`          // 海报URL地址
			FilmTypes    string      `json:"filmTypes"`    // 影片类型
			LikeNum      int         `json:"likeNum"`      // 想看人数
			Duration     int64       `json:"duration"`     // 时长，分钟
			Cast         string      `json:"cast"`         // 主演
			FilmId       int         `json:"filmId"`       // 影片id
			Grade        interface{} `json:"grade"`        // 评分
			Intro        string      `json:"intro"`        // 简介
			Name         string      `json:"name"`         // 影片名
		} `json:"list"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetSoonListResult struct {
	Result GetSoonListResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newGetSoonListResult(result GetSoonListResponse, body []byte, http gorequest.Response) *GetSoonListResult {
	return &GetSoonListResult{Result: result, Body: body, Http: http}
}

// GetSoonList 即将上映 https://www.showdoc.com.cn/1154868044931571/5866125707634369
func (c *Client) GetSoonList(ctx context.Context, cityId int, notMustParams ...gorequest.Params) (*GetSoonListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cityId", cityId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/movieapi/movie-info/get-soon-list", params)
	if err != nil {
		return newGetSoonListResult(GetSoonListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GetSoonListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGetSoonListResult(response, request.ResponseBody, request), err
}
