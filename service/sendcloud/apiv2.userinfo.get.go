package sendcloud

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ApiV2UserinfoGetResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID                 int64   `json:"id"`                   // 订单号
		ProductID          int     `json:"product_id"`           // 商品编号
		ProductName        string  `json:"product_name"`         // 商品名称
		ProductType        int     `json:"product_type"`         // 商品类型（1：充值，2：卡密，3：卡券，4：人工）
		ProductPrice       string  `json:"product_price"`        // 售价
		Quantity           int     `json:"quantity"`             // 购买数量
		TotalPrice         string  `json:"total_price"`          // 总支付价格
		RefundedAmount     float64 `json:"refunded_amount"`      // 已退款金额
		BuyerCustomerID    int     `json:"buyer_customer_id"`    // 买家编号
		BuyerCustomerName  string  `json:"buyer_customer_name"`  // 买家名称
		SellerCustomerID   int     `json:"seller_customer_id"`   // 卖家编号
		SellerCustomerName string  `json:"seller_customer_name"` // 卖家名称
		State              int     `json:"state"`                // 订单状态（100：等待发货，101：正在充值，200：交易成功，500：交易失败，501：未知状态）
		CreatedAt          string  `json:"created_at"`           // 下单时间
		RechargeAccount    string  `json:"recharge_account"`     // 充值账号
		ProgressInit       int     `json:"progress_init"`        // 充值进度：初始值
		ProgressNow        int     `json:"progress_now"`         // 充值进度：现在值
		ProgressTarget     int     `json:"progress_target"`      // 充值进度：目标值
		RechargeInfo       string  `json:"recharge_info"`        // 返回信息
		RechargeUrl        string  `json:"recharge_url"`         // 卡密充值网址
		Cards              []struct {
			No       string `json:"no"`
			Password string `json:"password"`
		} `json:"cards"` //【卡密类订单】卡密
		RechargeParams          string `json:"recharge_params"`                     //【充值类订单】
		OuterApiV2UserinfoGetID string `json:"outer_ApiV2UserinfoGet_id,omitempty"` // 外部订单号
	} `json:"data"`
}

type ApiV2UserinfoGetResult struct {
	Result ApiV2UserinfoGetResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newApiV2UserinfoGetResult(result ApiV2UserinfoGetResponse, body []byte, http gorequest.Response) *ApiV2UserinfoGetResult {
	return &ApiV2UserinfoGetResult{Result: result, Body: body, Http: http}
}

// ApiV2UserinfoGet 获取单个订单信息。
// 仅能获取自己购买的订单。
// http://doc.cqmeihu.cn/sales/ApiV2UserinfoGet-info.html
func (c *Client) ApiV2UserinfoGet(ctx context.Context, notMustParams ...gorequest.Params) (*ApiV2UserinfoGetResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("apiUser", c.GetApiUser())
	params.Set("apiKey", c.GetApiKey())
	// 请求
	request, err := c.request(ctx, apiUrl+"/apiv2/userinfo/get", params, http.MethodGet)
	if err != nil {
		return newApiV2UserinfoGetResult(ApiV2UserinfoGetResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiV2UserinfoGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiV2UserinfoGetResult(response, request.ResponseBody, request), err
}
