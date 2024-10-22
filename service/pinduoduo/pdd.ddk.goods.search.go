package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type GoodsSearchResponse struct {
	GoodsSearchResponse struct {
		GoodsList []struct {
			ActivityPromotionRate       int64    `json:"activity_promotion_rate"`
			ActivityTags                []int64  `json:"activity_tags"`
			ActivityType                int64    `json:"activity_type"`
			BrandName                   string   `json:"brand_name"`
			CashGiftAmount              int64    `json:"cash_gift_amount"`
			CatIds                      []int64  `json:"cat_ids"`
			CltCpnBatchSn               string   `json:"clt_cpn_batch_sn"`
			CltCpnDiscount              int64    `json:"clt_cpn_discount"`
			CltCpnEndTime               int64    `json:"clt_cpn_end_time"`
			CltCpnMinAmt                int64    `json:"clt_cpn_min_amt"`
			CltCpnQuantity              int64    `json:"clt_cpn_quantity"`
			CltCpnRemainQuantity        int64    `json:"clt_cpn_remain_quantity"`
			CltCpnStartTime             int64    `json:"clt_cpn_start_time"`
			CouponDiscount              int64    `json:"coupon_discount"`
			CouponEndTime               int64    `json:"coupon_end_time"`
			CouponMinOrderAmount        int64    `json:"coupon_min_order_amount"`
			CouponRemainQuantity        int64    `json:"coupon_remain_quantity"`
			CouponStartTime             int64    `json:"coupon_start_time"`
			CouponTotalQuantity         int64    `json:"coupon_total_quantity"`
			CreateAt                    int64    `json:"create_at"`
			DescTxt                     string   `json:"desc_txt"`
			ExtraCouponAmount           int64    `json:"extra_coupon_amount"`
			GoodsDesc                   string   `json:"goods_desc"`
			GoodsImageUrl               string   `json:"goods_image_url"`
			GoodsLabels                 []int64  `json:"goods_labels"`
			GoodsName                   string   `json:"goods_name"`
			GoodsSign                   string   `json:"goods_sign"`
			GoodsThumbnailUrl           string   `json:"goods_thumbnail_url"`
			HasCoupon                   bool     `json:"has_coupon"`
			HasMallCoupon               bool     `json:"has_mall_coupon"`
			HasMaterial                 bool     `json:"has_material"`
			LgstTxt                     string   `json:"lgst_txt"`
			MallCouponDiscountPct       int64    `json:"mall_coupon_discount_pct"`
			MallCouponEndTime           int64    `json:"mall_coupon_end_time"`
			MallCouponId                int64    `json:"mall_coupon_id"`
			MallCouponMaxDiscountAmount int64    `json:"mall_coupon_max_discount_amount"`
			MallCouponMinOrderAmount    int64    `json:"mall_coupon_min_order_amount"`
			MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity"`
			MallCouponStartTime         int64    `json:"mall_coupon_start_time"`
			MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity"`
			MallCps                     int64    `json:"mall_cps"`
			MallId                      int64    `json:"mall_id"`
			MallName                    string   `json:"mall_name"`
			MerchantType                int64    `json:"merchant_type"`
			MinGroupPrice               int64    `json:"min_group_price"`
			MinNormalPrice              int64    `json:"min_normal_price"`
			OnlySceneAuth               bool     `json:"only_scene_auth"`
			OptId                       int64    `json:"opt_id"`
			OptIds                      []int64  `json:"opt_ids"`
			OptName                     string   `json:"opt_name"`
			PlanType                    int64    `json:"plan_type"`
			PredictPromotionRate        int64    `json:"predict_promotion_rate"`
			PromotionRate               int64    `json:"promotion_rate"`
			SalesTip                    string   `json:"sales_tip"`
			SearchId                    string   `json:"search_id"`
			ServTxt                     string   `json:"serv_txt"`
			ServiceTags                 []int64  `json:"service_tags"`
			ShareRate                   int64    `json:"share_rate"`
			SubsidyAmount               int64    `json:"subsidy_amount"`
			SubsidyDuoAmountTenMillion  int64    `json:"subsidy_duo_amount_ten_million"`
			UnifiedTags                 []string `json:"unified_tags"`
			ZsDuoId                     int64    `json:"zs_duo_id"`
			GoodsId                     int64    `json:"goods_id"`
			CategoryId                  int64    `json:"category_id"`
			CategoryName                string   `json:"category_name"`
		} `json:"goods_list"`
		ListId     string `json:"list_id"`
		SearchId   string `json:"search_id"`
		TotalCount int64  `json:"total_count"`
	} `json:"goods_search_response"`
}

type GoodsSearchResult struct {
	Result GoodsSearchResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newGoodsSearchResult(result GoodsSearchResponse, body []byte, http gorequest.Response) *GoodsSearchResult {
	return &GoodsSearchResult{Result: result, Body: body, Http: http}
}

// GoodsSearch 多多进宝商品查询
// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.search
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.search
func (c *Client) GoodsSearch(ctx context.Context, notMustParams ...gorequest.Params) (*GoodsSearchResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.goods.search")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.goods.search", notMustParams...)
	params.Set("pid", c.GetPid())

	// 请求
	var response GoodsSearchResponse
	request, err := c.request(ctx, span, params, &response)
	return newGoodsSearchResult(response, request.ResponseBody, request), err
}
