package wechatunion

import (
	"errors"
)

type OrderSearch struct {
	Page                       int    `json:"page,omitempty"`                       // 页码，起始为 1
	PageSize                   int    `json:"pageSize,omitempty"`                   // 分页大小，最大 200
	StartTimestamp             string `json:"startTimestamp,omitempty"`             // 起始时间戳，单位为秒
	EndTimestamp               string `json:"endTimestamp,omitempty"`               // 结束时间戳，单位为秒
	CommissionStatus           string `json:"commissionStatus,omitempty"`           // 分佣状态
	SortByCommissionUpdateTime string `json:"sortByCommissionUpdateTime,omitempty"` // 是否按照分佣状态更新时间排序和筛选订单，1：是，0：否
	StartCommissionUpdateTime  string `json:"startCommissionUpdateTime,omitempty"`  // 分佣状态更新时间起始时间戳，单位为秒
	EndCommissionUpdateTime    string `json:"endCommissionUpdateTime,omitempty"`    // 分佣状态更新时间结束时间戳，单位为秒
}

type OrderSearchResult struct {
	Errcode   int    `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	OrderList []struct {
		OrderId            string `json:"orderId"`            // 订单ID
		PayTime            int    `json:"payTime"`            // 支付时间戳，单位为s
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
			Ratio                      int    `json:"ratio"`                      // 分佣比例，单位为万分之一
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
			CustomizeInfo string `json:"customizeInfo"` // 自定义信息
		} `json:"productList"` // 商品列表
	} `json:"orderList"` // 订单列表
	PageSize int `json:"pageSize"` // 分页大小
	TotalNum int `json:"totalNum"` // 订单总数
}

// OrderSearch 根据订单支付时间、订单分佣状态拉取订单详情 https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/order/order-info.html
func (c *Client) OrderSearch(notMustParams ...Params) (result OrderSearchResult, err error) {
	if len(c.getAccessToken()) <= 0 {
		return result, errors.New("调用凭证异常")
	}

	// 参数
	//params := c.NewParamsWith(notMustParams...)

	//if len(orderIdList) <= 0 || len(orderIdList) > 200 {
	//	return result, errors.New("未传入 orderIdList 或 orderIdList 超过上限 200")
	//}

	//body, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/order/info?access_token=%s", c.accessToken), map[string]interface{}{
	//	"orderIdList": orderIdList,
	//}, http.MethodPost)
	//if err != nil {
	//	return result, err
	//}
	//err = json.Unmarshal(body, &result)
	//if err != nil {
	//	return result, err
	//}
	//return result, err
	return
}
