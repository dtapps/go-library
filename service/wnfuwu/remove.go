package wnfuwu

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RemoveResponse struct {
	Errno  int64    `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string   `json:"errmsg"` // 错误描述
	Data   struct{} `json:"data"`
}

type RemoveResult struct {
	Result RemoveResponse     // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newRemoveResult(result RemoveResponse, body []byte, http gorequest.Response) *RemoveResult {
	return &RemoveResult{Result: result, Body: body, Http: http}
}

// Remove 申请撤单【已正式上线】
// https://www.showdoc.com.cn/dyr/9745453200292104
func (c *Client) Remove(ctx context.Context, outTradeNums string, notMustParams ...*gorequest.Params) (*RemoveResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId())
	if outTradeNums != "" {
		params.Set("out_trade_nums", outTradeNums)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/index/remove", params)
	if err != nil {
		return newRemoveResult(RemoveResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RemoveResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRemoveResult(response, request.ResponseBody, request), err
}
