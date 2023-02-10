package gddata

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GdpIndexResponse struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

type GdpIndexResult struct {
	Result GdpIndexResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newGdpIndexResult(result GdpIndexResponse, body []byte, http gorequest.Response, err error) *GdpIndexResult {
	return &GdpIndexResult{Result: result, Body: body, Http: http, Err: err}
}

// GdpIndex 地区生产总值指数（1978＝100）接口
// https://gddata.gd.gov.cn/opdata/index?chooseValue=apiForm&id=29000%2F03600017&sourceType
func (c *Client) GdpIndex(ctx context.Context, year string) *GdpIndexResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("year", year) // 统计年份
	// 请求
	request, err := c.request(ctx, apiUrl+fmt.Sprintf("MjkwMDBfMDM2MDAwMTc=?token=%s", c.GetToken()), params, http.GET)
	// 定义
	var response GdpIndexResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGdpIndexResult(response, request.ResponseBody, request, err)
}
