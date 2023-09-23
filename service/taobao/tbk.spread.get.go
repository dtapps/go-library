package taobao

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TbkSpreadGetResponse struct {
	TbkSpreadGetResponse struct {
		Results struct {
			TbkSpread []struct {
				Content string `json:"content"`
				ErrMsg  string `json:"err_msg"`
			} `json:"tbk_spread"`
		} `json:"results"`
		TotalResults int    `json:"total_results"`
		RequestId    string `json:"request_id"`
	} `json:"tbk_spread_get_response"`
}

type TbkSpreadGetResult struct {
	Result TbkSpreadGetResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newTbkSpreadGetResult(result TbkSpreadGetResponse, body []byte, http gorequest.Response) *TbkSpreadGetResult {
	return &TbkSpreadGetResult{Result: result, Body: body, Http: http}
}

// TbkSpreadGet 淘宝客-公用-长链转短链
// https://open.taobao.com/api.htm?docId=27832&docType=2&source=search
func (c *Client) TbkSpreadGet(ctx context.Context, notMustParams ...*gorequest.Params) (*TbkSpreadGetResult, error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.spread.get", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newTbkSpreadGetResult(TbkSpreadGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkSpreadGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkSpreadGetResult(response, request.ResponseBody, request), err
}
