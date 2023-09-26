package taobao

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TbkDgNewuserOrderGetResponse struct {
	TbkDgNewuserOrderGetResponse struct {
		Data struct {
			Model          string `json:"model"`
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
		RequestId string `json:"request_id"`
	} `json:"tbk_tpwd_create_response"`
}

type TbkDgNewuserOrderGetResult struct {
	Result TbkDgNewuserOrderGetResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newTbkDgNewuserOrderGetResult(result TbkDgNewuserOrderGetResponse, body []byte, http gorequest.Response) *TbkDgNewuserOrderGetResult {
	return &TbkDgNewuserOrderGetResult{Result: result, Body: body, Http: http}
}

// TbkDgNewuserOrderGet 淘宝客-推广者-新用户订单明细查询
// https://open.taobao.com/api.htm?docId=33892&docType=2
func (c *Client) TbkDgNewuserOrderGet(ctx context.Context, notMustParams ...gorequest.Params) (*TbkDgNewuserOrderGetResult, error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.dg.newuser.order.get", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newTbkDgNewuserOrderGetResult(TbkDgNewuserOrderGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TbkDgNewuserOrderGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTbkDgNewuserOrderGetResult(response, request.ResponseBody, request), err
}
