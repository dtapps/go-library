package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestPowerQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		Id             uint   `json:"id,omitempty"`
		Fanli          string `json:"fanli"`            // 平台返利金额
		Amount         int64  `json:"amount"`           // 充值金额
		CostPrice      string `json:"cost_price"`       // 成本价格
		Status         int    `json:"status"`           // 订单状态 0 待支付 1 已付 充值中 2充值成功 3充值失败 需要退款 4退款成功 5已超时 6待充值 7 已匹配 8已存单 9 已取消 10返销 11部分到账 12取消中
		OrderNumber    string `json:"order_number"`     // 平台单号
		OrderNo        string `json:"order_no"`         // 第三方单号
		OrgOrderNumber string `json:"org_order_number"` // 组织订单号
		CardId         string `json:"card_id"`          // 充值卡ID
		ArrivedAmount  int64  `json:"arrived_amount"`   // 到账金额
		Reason         string `json:"reason,omitempty"` // 失败原因
	} `json:"data"`
}

type RestPowerQueryResult struct {
	Result RestPowerQueryResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newRestPowerQueryResult(result RestPowerQueryResponse, body []byte, http gorequest.Response) *RestPowerQueryResult {
	return &RestPowerQueryResult{Result: result, Body: body, Http: http}
}

// RestPowerQuery 电费订单查询
// https://open.wikeyun.cn/#/apiDocument/9/document/313
func (c *Client) RestPowerQuery(ctx context.Context, notMustParams ...gorequest.Params) (*RestPowerQueryResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/query", params)
	if err != nil {
		return newRestPowerQueryResult(RestPowerQueryResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerQueryResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerQueryResult(response, request.ResponseBody, request), err
}

func (resp RestPowerQueryResponse) GetStatusDesc(status int) string {
	switch status {
	case 1:
		return "充值中"
	case 2:
		return "充值成功"
	case 3:
		return "充值失败"
	case 4:
		return "退款成功"
	case 5:
		return "已超时"
	case 6:
		return "待充值"
	case 7:
		return "已匹配"
	case 8:
		return "已存单"
	case 9:
		return "已取消"
	case 10:
		return "返销"
	case 11:
		return "部分到账"
	case 12:
		return "取消中"
	}
	return "待支付"
}
