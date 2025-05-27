package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type OauthCashGiftDataQueryList struct {
	Amount        float64 `json:"amount"`         // 礼金券创建总金额，单位为分
	CashGiftId    float64 `json:"cash_gift_id"`   // 礼金ID
	CashGiftName  string  `json:"cash_gift_name"` // 礼金名称
	CouponAmount  float64 `json:"couponAmount"`   // 礼金券面额
	FetchAmount   float64 `json:"fetch_amount"`   // 已领取礼金券总金额，单位为分（实时数据）
	FetchQuantity int64   `json:"fetch_quantity"` // 已领取礼金券数量（实时数据）
	GoodsInfoList []struct {
		CouponDiscount float64 `json:"coupon_discount"` // 商品优惠券面额，单位为分
		GoodsName      string  `json:"goods_name"`      // 商品名称
		GoodsPrice     float64 `json:"goods_price"`     // 商品原价，单位为分
		GoodsSign      string  `json:"goods_sign"`      // 商品goodsSign，支持通过goodsSign查询商品。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
		Rate           int64   `json:"rate"`            // 商品佣金比例，千分比
	} `json:"goods_info_list"` // 商品列表信息
	OrderCouponAmount  float64 `json:"order_coupon_amount"`  // 礼金订单使用的券总金额，单位为分（实时数据）
	OrderGmv           float64 `json:"order_gmv"`            // 礼金订单产生的总GMV，单位为分（实时数据）
	OrderQuantity      int64   `json:"order_quantity"`       // 礼金订单数量（实时数据）
	PrePromotionAmount float64 `json:"pre_promotion_amount"` // 礼金订单预估佣金，单位为分（实时数据）
	Quantity           int64   `json:"quantity"`             // 礼金券创建总数量
	RefundAmount       float64 `json:"refund_amount"`        // 退回礼金券总金额，单位为分
	RefundQuantity     int64   `json:"refund_quantity"`      // 退回礼金券数量
	Status             int64   `json:"status"`               // 礼金状态：1-未生效；2-生效中；3-已过期；4-活动中止（用户主动停止）；5-活动中止（佣金降低）；6-活动中止（推广活动异常）
}

type OauthCashGiftDataQuery struct {
	CashgiftDataResponse struct {
		AvailableBalance float64                      `json:"available_balance"` // 礼金账户余额，单位为分
		List             []OauthCashGiftDataQueryList `json:"list"`              // 多多礼金数据列表
		Total            int64                        `json:"total"`             // 请求到的结果数
	} `json:"cashgift_data_response"`
}

// OauthCashGiftDataQuery 查询多多礼金效果数据
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.data.query
func (c *Client) OauthCashGiftDataQuery(ctx context.Context, notMustParams ...*gorequest.Params) (response OauthCashGiftDataQuery, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.data.query", notMustParams...)

	// 请求
	err = c.request(ctx, params, &response)
	return
}
