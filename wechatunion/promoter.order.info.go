package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type PromoterOrderInfoResponse struct {
	Errcode   int    `json:"errcode"` // 错误码
	Errmsg    string `json:"errmsg"`  // 错误信息
	OrderList []struct {
		OrderId            string `json:"orderId"`            // 订单ID
		PayTime            int64  `json:"payTime"`            // 支付时间戳，单位为s
		ConfirmReceiptTime int    `json:"confirmReceiptTime"` // 确认收货时间戳，单位为s，没有时为0
		ShopName           string `json:"shopName"`           // 店铺名称
		ShopAppid          string `json:"shopAppid"`          // 店铺 Appid
		ProductList        []struct {
			ProductId                  string `json:"productId"`                  // 商品SPU ID
			SkuId                      string `json:"skuId"`                      // sku ID
			Title                      string `json:"title"`                      // 商品名称
			ThumbImg                   string `json:"thumbImg"`                   // 商品缩略图 url
			Price                      string `json:"price"`                      // 商品成交总价，前带单位 ¥
			ProductCnt                 int    `json:"productCnt"`                 // 成交数量
			Ratio                      int64  `json:"ratio"`                      // 分佣比例，单位为万分之一
			CommissionStatus           string `json:"commissionStatus"`           // 分佣状态
			CommissionStatusUpdateTime string `json:"commissionStatusUpdateTime"` // 分佣状态更新时间戳，单位为s
			ProfitShardingSucTime      string `json:"profitShardingSucTime"`      // 结算时间，当分佣状态为已结算才有值，单位为s
			Commission                 string `json:"commission"`                 // 分佣金额，前带单位 ¥
			EstimatedCommission        int    `json:"estimatedCommission"`        // 预估分佣金额，单位为分
			CategoryStr                string `json:"categoryStr"`                // 类目名称，多个用英文逗号分隔
			PromotionInfo              struct {
				PromotionSourcePid  string `json:"promotionSourcePid"`  // 推广位 id
				PromotionSourceName string `json:"promotionSourceName"` // 推广位名称
			} `json:"promotionInfo"` // 推广信息
			CustomizeInfo string `json:"customizeInfo"` // 	自定义信息
		} `json:"productList"` // 商品列表
		CustomUserId string `json:"customUserId"` // 自定义用户参数
		UserNickName string `json:"userNickName"` // 用户昵称
		OrderPrice   string `json:"orderPrice"`   // 支付金额，单位为分
	} `json:"orderList"`
}

type PromoterOrderInfoResult struct {
	Result PromoterOrderInfoResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func NewPromoterOrderInfoResult(result PromoterOrderInfoResponse, body []byte, http gorequest.Response, err error) *PromoterOrderInfoResult {
	return &PromoterOrderInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// PromoterOrderInfo 根据订单ID查询订单详情
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/order/order-info.html#_1-%E6%A0%B9%E6%8D%AE%E8%AE%A2%E5%8D%95ID%E6%9F%A5%E8%AF%A2%E8%AE%A2%E5%8D%95%E8%AF%A6%E6%83%85
func (app *App) PromoterOrderInfo(orderId ...string) *PromoterOrderInfoResult {
	app.accessToken = app.GetAccessToken()
	// 参数
	params := app.NewParamsWith()
	var orderIdList []any
	for _, v := range orderId {
		orderIdList = append(orderIdList, v)
	}
	params.Set("orderIdList", orderIdList)
	// 请求
	request, err := app.request(UnionUrl+fmt.Sprintf("/promoter/order/info?access_token=%s", app.accessToken), params, http.MethodPost)
	// 定义
	var response PromoterOrderInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewPromoterOrderInfoResult(response, request.ResponseBody, request, err)
}
