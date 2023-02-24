package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type billSubMerchantFundFlowBillGetResponse struct {
	DownloadBillCount int `json:"download_bill_count"` // 下载信息总数
	DownloadBillList  []struct {
		BillSequence int    `json:"bill_sequence"` // 账单文件序号
		DownloadUrl  string `json:"download_url"`  // 下载地址
		EncryptKey   string `json:"encrypt_key"`   // 加密密钥
		HashType     string `json:"hash_type"`     // 哈希类型
		HashValue    string `json:"hash_value"`    // 哈希值
		Nonce        string `json:"nonce"`         // 随机字符串
	} `json:"download_bill_list"` // 下载信息明细
}

type billSubMerchantFundFlowBillGetResult struct {
	Result   billSubMerchantFundFlowBillGetResponse // 结果
	Body     []byte                                 // 内容
	Http     gorequest.Response                     // 请求
	Err      error                                  // 错误
	ApiError ApiError                               // 接口错误
}

func newbillSubMerchantFundFlowBillGetResult(result billSubMerchantFundFlowBillGetResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *billSubMerchantFundFlowBillGetResult {
	return &billSubMerchantFundFlowBillGetResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// billSubMerchantFundFlowBillGet 申请单个子商户资金账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_12.shtml
func (c *Client) billSubMerchantFundFlowBillGet(ctx context.Context, notMustParams ...gorequest.Params) *billSubMerchantFundFlowBillGetResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/bill/sub-merchant-fundflowbill", params, http.MethodGet)
	if err != nil {
		return newbillSubMerchantFundFlowBillGetResult(billSubMerchantFundFlowBillGetResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response billSubMerchantFundFlowBillGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newbillSubMerchantFundFlowBillGetResult(response, request.ResponseBody, request, err, apiError)
}
