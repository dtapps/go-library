package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type OrderDetailGetResponse struct {
	OrderDetailResponse struct {
		ActivityTags                      []int64 `json:"activity_tags"`                                     // 商品活动标记数组，例：[4,7]，4-秒杀 7-百亿补贴等
		AuthDuoID                         int64   `json:"auth_duo_id"`                                       // 多多客工具id
		BatchNo                           string  `json:"batch_no"`                                          // 结算批次号
		CashGiftID                        int64   `json:"cash_gift_id,omitempty"`                            // 订单关联礼金活动Id
		CatIds                            []any   `json:"cat_ids"`                                           // 商品一~四级类目ID列表
		CpaNew                            int64   `json:"cpa_new"`                                           // 是否是 cpa 新用户，1表示是，0表示否
		CustomParameters                  string  `json:"custom_parameters"`                                 // 自定义参数
		FailReason                        string  `json:"fail_reason,omitempty"`                             // 订单审核失败/惩罚原因
		GoodsCategoryName                 string  `json:"goods_category_name"`                               // 商品一级类目名称
		GoodsID                           int64   `json:"goods_id"`                                          // 商品ID
		GoodsName                         string  `json:"goods_name"`                                        // 商品标题
		GoodsPrice                        int64   `json:"goods_price"`                                       // 订单中sku的单件价格，单位为分
		GoodsQuantity                     int64   `json:"goods_quantity"`                                    // 购买商品的数量
		GoodsSign                         string  `json:"goods_sign"`                                        // goodsSign是加密后的goodsId，goodsId已下线，请使用goodsSign来替代。需要注意的是：推广链接带有goodsSign信息时，订单会返回原goodsSign；反之，会生成新的goodsSign返回。
		GoodsThumbnailUrl                 string  `json:"goods_thumbnail_url"`                               // 商品缩略图
		GroupID                           int64   `json:"group_id"`                                          // 成团编号
		InTenMillionSubsidyQuota          int64   `json:"in_ten_million_subsidy_quota,omitempty"`            // 计入千万补贴额度(仅top渠道享受) 值为0时不计入 其他为null
		IsDirect                          int64   `json:"is_direct"`                                         // 是否直推 ，1表示是，0表示否
		MallID                            int64   `json:"mall_id"`                                           // 店铺id
		MallName                          string  `json:"mall_name"`                                         // 店铺名称
		NoSubsidyReason                   string  `json:"no_subsidy_reason,omitempty"`                       // 非补贴订单原因，例如："商品补贴达上限"，"达到单个用户下单上限"，"非指定落地页直推订单"，"订单超过2个月未审核成功"等
		NotInTenMillionSubsidyQuotaReason string  `json:"not_in_ten_million_subsidy_quota_reason,omitempty"` // 不计入千万补贴额度原因
		OrderAmount                       int64   `json:"order_amount"`                                      // 实际支付金额，单位为分
		OrderCreateTime                   int64   `json:"order_create_time"`                                 // 订单生成时间，UNIX时间戳
		OrderGroupSuccessTime             int64   `json:"order_group_success_time"`                          // 成团时间
		OrderID                           string  `json:"order_id"`                                          // 订单ID
		OrderModifyAt                     int64   `json:"order_modify_at"`                                   // 最后更新时间
		OrderPayTime                      int64   `json:"order_pay_time"`                                    // 支付时间
		OrderReceiveTime                  int64   `json:"order_receive_time,omitempty"`                      // 确认收货时间
		OrderSettleTime                   int64   `json:"order_settle_time,omitempty"`                       // 结算时间
		OrderSn                           string  `json:"order_sn"`                                          // 推广订单编号
		OrderStatus                       int64   `json:"order_status"`                                      // 订单状态：0-已支付；1-已成团；2-确认收货；3-审核成功；4-审核失败（不可提现）；5-已经结算 ;10-已处罚
		OrderStatusDesc                   string  `json:"order_status_desc"`                                 // 订单状态描述
		OrderVerifyTime                   int64   `json:"order_verify_time,omitempty"`                       // 审核时间
		PID                               string  `json:"p_id"`                                              // 推广位ID
		PlatformDiscount                  int64   `json:"platform_discount"`                                 // 平台券金额，表示该订单使用的平台券金额，单位分
		PriceCompareStatus                int64   `json:"price_compare_status"`                              // 比价状态：0：正常，1：比价
		PromotionAmount                   int64   `json:"promotion_amount"`                                  // 佣金金额，单位为分
		PromotionRate                     int64   `json:"promotion_rate"`                                    // 佣金比例，千分比
		RedPacketType                     int64   `json:"red_packet_type"`                                   // 超级红包补贴类型：0-非红包补贴订单，1-季度新用户补贴
		SepDuoID                          int64   `json:"sep_duo_id"`                                        // 直播间订单推广duoId
		SepMarketFee                      int64   `json:"sep_market_fee"`                                    // 直播间推广佣金
		SepParameters                     string  `json:"sep_parameters"`                                    // 直播间推广自定义参数
		SepPID                            string  `json:"sep_pid"`                                           // 直播间订单推广位
		SepRate                           int64   `json:"sep_rate"`                                          // 直播间推广佣金比例
		ShareAmount                       int64   `json:"share_amount"`                                      // 招商分成服务费金额，单位为分
		ShareRate                         int64   `json:"share_rate,omitempty"`                              // 招商分成服务费比例，千分比
		SubsidyAmount                     int64   `json:"subsidy_amount,omitempty"`                          // 优势渠道专属商品补贴金额，单位为分。针对优质渠道的补贴活动，指定优势渠道可通过推广该商品获取相应补贴。补贴活动入口：[进宝网站-官方活动]
		SubsidyDuoAmountLevel             int64   `json:"subsidy_duo_amount_level"`                          // 等级补贴给渠道的收入补贴，不允许直接给下级代理展示，单位为分
		SubsidyDuoAmountTenMillion        int64   `json:"subsidy_duo_amount_ten_million,omitempty"`          // 官方活动给渠道的收入补贴金额，不允许直接给下级代理展示，单位为分
		SubsidyOrderRemark                string  `json:"subsidy_order_remark,omitempty"`                    // 补贴订单备注
		SubsidyType                       int64   `json:"subsidy_type"`                                      // 订单补贴类型：0-非补贴订单，1-千万补贴，2-社群补贴，3-多多星选，4-品牌优选，5-千万神券，8-拼团享多多
		Type                              int64   `json:"type"`                                              // 下单场景类型：0-单品推广，1-红包活动推广，4-多多进宝商城推广，7-今日爆款，8-品牌清仓，9-1.9包邮，77-刮刮卡活动推广，94-充值中心，101-品牌黑卡，103-百亿补贴频道，104-内购清单频道，105-超级红包
		ZsDuoID                           int64   `json:"zs_duo_id"`                                         // 招商多多客id
		BandanRiskConsult                 int64   `json:"bandan_risk_consult"`                               // 预判断是否为代购订单，-1（默认）表示未出结果，0表示预判不是代购订单，1表示代购订单，具体请以最后审核状态为准
		ActivityType                      int64   `json:"activity_type"`                                     // 活动类型
	} `json:"order_detail_response"`
}

type OrderDetailGetResult struct {
	Result OrderDetailGetResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newOrderDetailGetResult(result OrderDetailGetResponse, body []byte, http gorequest.Response) *OrderDetailGetResult {
	return &OrderDetailGetResult{Result: result, Body: body, Http: http}
}

// OrderDetailGet 多多进宝商品查询 https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.order.detail.get
func (c *Client) OrderDetailGet(ctx context.Context, orderSn string, notMustParams ...gorequest.Params) (*OrderDetailGetResult, error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.order.detail.get", notMustParams...)
	params.Set("order_sn", orderSn)

	// 请求
	var response OrderDetailGetResponse
	request, err := c.request(ctx, params, &response)
	return newOrderDetailGetResult(response, request.ResponseBody, request), err
}
