package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type GoodsRecommendGetList struct {
	ActivityPromotionRate      int64    `json:"activity_promotion_rate"`
	ActivityTags               []int64  `json:"activity_tags"`
	BrandName                  string   `json:"brand_name"`
	CashGiftAmount             int64    `json:"cash_gift_amount"`
	CatId                      string   `json:"cat_id"`
	CatIds                     []int64  `json:"cat_ids"`
	CouponDiscount             int64    `json:"coupon_discount"`
	CouponEndTime              int64    `json:"coupon_end_time"`
	CouponMinOrderAmount       int64    `json:"coupon_min_order_amount"`
	CouponPrice                int64    `json:"coupon_price"`
	CouponRemainQuantity       int64    `json:"coupon_remain_quantity"`
	CouponStartTime            int64    `json:"coupon_start_time"`
	CouponTotalQuantity        int64    `json:"coupon_total_quantity"`
	CreateAt                   int64    `json:"create_at"`
	DescTxt                    string   `json:"desc_txt"`
	ExtraCouponAmount          int64    `json:"extra_coupon_amount"`
	GoodsDesc                  string   `json:"goods_desc"`
	GoodsImageUrl              string   `json:"goods_image_url"`
	GoodsLabels                []int64  `json:"goods_labels"`
	GoodsName                  string   `json:"goods_name"`
	GoodsRate                  int64    `json:"goods_rate"`
	GoodsSign                  string   `json:"goods_sign"`
	GoodsThumbnailUrl          string   `json:"goods_thumbnail_url"`
	GoodsType                  int64    `json:"goods_type"`
	HasCoupon                  bool     `json:"has_coupon"`
	HasMaterial                bool     `json:"has_material"`
	LgstTxt                    string   `json:"lgst_txt"`
	MallId                     int64    `json:"mall_id"`
	MallName                   string   `json:"mall_name"`
	MarketFee                  int64    `json:"market_fee"`
	MerchantType               string   `json:"merchant_type"`
	MinGroupPrice              int64    `json:"min_group_price"`
	MinNormalPrice             int64    `json:"min_normal_price"`
	OptId                      string   `json:"opt_id"`
	OptIds                     []int64  `json:"opt_ids"`
	OptName                    string   `json:"opt_name"`
	PredictPromotionRate       int64    `json:"predict_promotion_rate"`
	PromotionRate              int64    `json:"promotion_rate"`
	QrCodeImageUrl             string   `json:"qr_code_image_url"`
	RealtimeSalesTip           string   `json:"realtime_sales_tip"`
	SalesTip                   string   `json:"sales_tip"`
	SearchId                   string   `json:"search_id"`
	ServTxt                    string   `json:"serv_txt"`
	ShareDesc                  string   `json:"share_desc"`
	ShareRate                  int64    `json:"share_rate"`
	SubsidyAmount              int64    `json:"subsidy_amount"`
	SubsidyDuoAmountTenMillion int64    `json:"subsidy_duo_amount_ten_million"`
	UnifiedTags                []string `json:"unified_tags"`
	GoodsId                    int64    `json:"goods_id"`
	CategoryId                 string   `json:"category_id"`
	CategoryName               string   `json:"category_name"`
}

type GoodsRecommendGet struct {
	GoodsBasicDetailResponse struct {
		List     []GoodsRecommendGetList `json:"list"`
		ListId   string                  `json:"list_id"`
		SearchId string                  `json:"search_id"`
		Total    int64                   `json:"total"`
	} `json:"goods_basic_detail_response"`
}

// GoodsRecommendGet 多多进宝商品推荐API
// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.recommend.get
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.recommend.get
func (c *Client) GoodsRecommendGet(ctx context.Context, notMustParams ...*gorequest.Params) (response GoodsRecommendGet, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.goods.recommend.get", notMustParams...)
	params.Set("pid", c.GetPid())

	// 请求
	err = c.request(ctx, params, &response)
	return
}
