package pinduoduo

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type GoodsSearchResponse struct {
	GoodsSearchResponse struct {
		GoodsList []struct {
			ActivityPromotionRate       int      `json:"activity_promotion_rate"`
			ActivityTags                []int    `json:"activity_tags"`
			ActivityType                int      `json:"activity_type"`
			BrandName                   string   `json:"brand_name"`
			CashGiftAmount              int      `json:"cash_gift_amount"`
			CatIds                      []int    `json:"cat_ids"`
			CltCpnBatchSn               string   `json:"clt_cpn_batch_sn"`
			CltCpnDiscount              int      `json:"clt_cpn_discount"`
			CltCpnEndTime               int      `json:"clt_cpn_end_time"`
			CltCpnMinAmt                int      `json:"clt_cpn_min_amt"`
			CltCpnQuantity              int      `json:"clt_cpn_quantity"`
			CltCpnRemainQuantity        int      `json:"clt_cpn_remain_quantity"`
			CltCpnStartTime             int      `json:"clt_cpn_start_time"`
			CouponDiscount              int64    `json:"coupon_discount"`
			CouponEndTime               int      `json:"coupon_end_time"`
			CouponMinOrderAmount        int      `json:"coupon_min_order_amount"`
			CouponRemainQuantity        int      `json:"coupon_remain_quantity"`
			CouponStartTime             int      `json:"coupon_start_time"`
			CouponTotalQuantity         int      `json:"coupon_total_quantity"`
			CreateAt                    int      `json:"create_at"`
			DescTxt                     string   `json:"desc_txt"`
			ExtraCouponAmount           int      `json:"extra_coupon_amount"`
			GoodsDesc                   string   `json:"goods_desc"`
			GoodsImageUrl               string   `json:"goods_image_url"`
			GoodsLabels                 []int    `json:"goods_labels"`
			GoodsName                   string   `json:"goods_name"`
			GoodsSign                   string   `json:"goods_sign"`
			GoodsThumbnailUrl           string   `json:"goods_thumbnail_url"`
			HasCoupon                   bool     `json:"has_coupon"`
			HasMallCoupon               bool     `json:"has_mall_coupon"`
			HasMaterial                 bool     `json:"has_material"`
			LgstTxt                     string   `json:"lgst_txt"`
			MallCouponDiscountPct       int      `json:"mall_coupon_discount_pct"`
			MallCouponEndTime           int      `json:"mall_coupon_end_time"`
			MallCouponId                int      `json:"mall_coupon_id"`
			MallCouponMaxDiscountAmount int      `json:"mall_coupon_max_discount_amount"`
			MallCouponMinOrderAmount    int      `json:"mall_coupon_min_order_amount"`
			MallCouponRemainQuantity    int      `json:"mall_coupon_remain_quantity"`
			MallCouponStartTime         int      `json:"mall_coupon_start_time"`
			MallCouponTotalQuantity     int      `json:"mall_coupon_total_quantity"`
			MallCps                     int      `json:"mall_cps"`
			MallId                      int64    `json:"mall_id"`
			MallName                    string   `json:"mall_name"`
			MerchantType                int      `json:"merchant_type"`
			MinGroupPrice               int64    `json:"min_group_price"`
			MinNormalPrice              int      `json:"min_normal_price"`
			OnlySceneAuth               bool     `json:"only_scene_auth"`
			OptId                       int      `json:"opt_id"`
			OptIds                      []int    `json:"opt_ids"`
			OptName                     string   `json:"opt_name"`
			PlanType                    int      `json:"plan_type"`
			PredictPromotionRate        int      `json:"predict_promotion_rate"`
			PromotionRate               int64    `json:"promotion_rate"`
			SalesTip                    string   `json:"sales_tip"`
			SearchId                    string   `json:"search_id"`
			ServTxt                     string   `json:"serv_txt"`
			ServiceTags                 []int    `json:"service_tags"`
			ShareRate                   int      `json:"share_rate"`
			SubsidyAmount               int      `json:"subsidy_amount"`
			SubsidyDuoAmountTenMillion  int      `json:"subsidy_duo_amount_ten_million"`
			UnifiedTags                 []string `json:"unified_tags"`
			ZsDuoId                     int      `json:"zs_duo_id"`
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
	Err    error               // 错误
}

func NewGoodsSearchResult(result GoodsSearchResponse, body []byte, http gorequest.Response, err error) *GoodsSearchResult {
	return &GoodsSearchResult{Result: result, Body: body, Http: http, Err: err}
}

// GoodsSearch 多多进宝商品查询
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.search
func (c *Client) GoodsSearch(notMustParams ...Params) *GoodsSearchResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.goods.search", notMustParams...)
	params.Set("pid", c.config.Pid)
	// 请求
	request, err := c.request(params)
	// 定义
	var response GoodsSearchResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewGoodsSearchResult(response, request.ResponseBody, request, err)
}
