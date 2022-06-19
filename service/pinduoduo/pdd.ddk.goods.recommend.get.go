package pinduoduo

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type GoodsRecommendGetResponse struct {
	GoodsBasicDetailResponse struct {
		List []struct {
			ActivityPromotionRate      int      `json:"activity_promotion_rate"`
			ActivityTags               []int    `json:"activity_tags"`
			BrandName                  string   `json:"brand_name"`
			CashGiftAmount             int      `json:"cash_gift_amount"`
			CatId                      string   `json:"cat_id"`
			CatIds                     []int    `json:"cat_ids"`
			CouponDiscount             int64    `json:"coupon_discount"`
			CouponEndTime              int      `json:"coupon_end_time"`
			CouponMinOrderAmount       int      `json:"coupon_min_order_amount"`
			CouponPrice                int      `json:"coupon_price"`
			CouponRemainQuantity       int      `json:"coupon_remain_quantity"`
			CouponStartTime            int      `json:"coupon_start_time"`
			CouponTotalQuantity        int      `json:"coupon_total_quantity"`
			CreateAt                   int      `json:"create_at"`
			DescTxt                    string   `json:"desc_txt"`
			ExtraCouponAmount          int      `json:"extra_coupon_amount"`
			GoodsDesc                  string   `json:"goods_desc"`
			GoodsImageUrl              string   `json:"goods_image_url"`
			GoodsLabels                []int    `json:"goods_labels"`
			GoodsName                  string   `json:"goods_name"`
			GoodsRate                  int      `json:"goods_rate"`
			GoodsSign                  string   `json:"goods_sign"`
			GoodsThumbnailUrl          string   `json:"goods_thumbnail_url"`
			GoodsType                  int      `json:"goods_type"`
			HasCoupon                  bool     `json:"has_coupon"`
			HasMaterial                bool     `json:"has_material"`
			LgstTxt                    string   `json:"lgst_txt"`
			MallId                     int64    `json:"mall_id"`
			MallName                   string   `json:"mall_name"`
			MarketFee                  int      `json:"market_fee"`
			MerchantType               string   `json:"merchant_type"`
			MinGroupPrice              int64    `json:"min_group_price"`
			MinNormalPrice             int      `json:"min_normal_price"`
			OptId                      string   `json:"opt_id"`
			OptIds                     []int    `json:"opt_ids"`
			OptName                    string   `json:"opt_name"`
			PredictPromotionRate       int      `json:"predict_promotion_rate"`
			PromotionRate              int64    `json:"promotion_rate"`
			QrCodeImageUrl             string   `json:"qr_code_image_url"`
			RealtimeSalesTip           string   `json:"realtime_sales_tip"`
			SalesTip                   string   `json:"sales_tip"`
			SearchId                   string   `json:"search_id"`
			ServTxt                    string   `json:"serv_txt"`
			ShareDesc                  string   `json:"share_desc"`
			ShareRate                  int      `json:"share_rate"`
			SubsidyAmount              int      `json:"subsidy_amount"`
			SubsidyDuoAmountTenMillion int      `json:"subsidy_duo_amount_ten_million"`
			UnifiedTags                []string `json:"unified_tags"`
			GoodsId                    int64    `json:"goods_id"`
			CategoryId                 string   `json:"category_id"`
			CategoryName               string   `json:"category_name"`
		} `json:"list"`
		ListId   string `json:"list_id"`
		SearchId string `json:"search_id"`
		Total    int64  `json:"total"`
	} `json:"goods_basic_detail_response"`
}

type GoodsRecommendGetResult struct {
	Result GoodsRecommendGetResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func newGoodsRecommendGetResult(result GoodsRecommendGetResponse, body []byte, http gorequest.Response, err error) *GoodsRecommendGetResult {
	return &GoodsRecommendGetResult{Result: result, Body: body, Http: http, Err: err}
}

// GoodsRecommendGet 多多进宝商品推荐API
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.recommend.get
func (c *Client) GoodsRecommendGet(notMustParams ...Params) *GoodsRecommendGetResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.goods.recommend.get", notMustParams...)
	params.Set("pid", c.config.Pid)
	// 请求
	request, err := c.request(params)
	// 定义
	var response GoodsRecommendGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newGoodsRecommendGetResult(response, request.ResponseBody, request, err)
}
