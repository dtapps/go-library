package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type PromoterOrderSearchResponse struct {
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
	} `json:"orderList"` // 订单列表
	PageSize int `json:"pageSize"` // 分页大小
	TotalNum int `json:"totalNum"` // 订单总数
}

type PromoterOrderSearchResult struct {
	Result PromoterOrderSearchResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func NewPromoterOrderSearchResult(result PromoterOrderSearchResponse, body []byte, http gorequest.Response, err error) *PromoterOrderSearchResult {
	return &PromoterOrderSearchResult{Result: result, Body: body, Http: http, Err: err}
}

// PromoterOrderSearch 根据订单支付时间、订单分佣状态拉取订单详情
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/order/order-info.html#_2-%E6%A0%B9%E6%8D%AE%E8%AE%A2%E5%8D%95%E6%94%AF%E4%BB%98%E6%97%B6%E9%97%B4%E3%80%81%E8%AE%A2%E5%8D%95%E5%88%86%E4%BD%A3%E7%8A%B6%E6%80%81%E6%8B%89%E5%8F%96%E8%AE%A2%E5%8D%95%E8%AF%A6%E6%83%85
func (app *App) PromoterOrderSearch(notMustParams ...Params) *PromoterOrderSearchResult {
	app.accessToken = app.GetAccessToken()
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request(UnionUrl+fmt.Sprintf("/promoter/order/search?access_token=%s", app.accessToken), params, http.MethodGet)
	// 定义
	var response PromoterOrderSearchResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewPromoterOrderSearchResult(response, request.ResponseBody, request, err)
}
