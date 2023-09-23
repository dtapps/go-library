package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type BillTradeBillGetResponse struct {
	DownloadUrl string `json:"download_url"` // 哈希类型
	HashType    string `json:"hash_type"`    // 哈希值
	HashValue   string `json:"hash_value"`   // 账单下载地址
}

type BillTradeBillGetResult struct {
	Result BillTradeBillGetResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newBillTradeBillGetResult(result BillTradeBillGetResponse, body []byte, http gorequest.Response) *BillTradeBillGetResult {
	return &BillTradeBillGetResult{Result: result, Body: body, Http: http}
}

// BillTradeBillGet 申请交易账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_6.shtml
func (c *Client) BillTradeBillGet(ctx context.Context, notMustParams ...*gorequest.Params) (*BillTradeBillGetResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/bill/tradebill", params, http.MethodGet)
	if err != nil {
		return newBillTradeBillGetResult(BillTradeBillGetResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response BillTradeBillGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newBillTradeBillGetResult(response, request.ResponseBody, request), apiError, err
}
