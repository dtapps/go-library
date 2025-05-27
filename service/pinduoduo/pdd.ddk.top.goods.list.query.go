package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type TopGoodsListQueryList struct {
	CatIds               []int64  `json:"cat_ids"`
	CouponDiscount       int64    `json:"coupon_discount"`
	CouponEndTime        int64    `json:"coupon_end_time"`
	CouponMinOrderAmount int64    `json:"coupon_min_order_amount"`
	CouponRemainQuantity int64    `json:"coupon_remain_quantity"`
	CouponStartTime      int64    `json:"coupon_start_time"`
	CouponTotalQuantity  int64    `json:"coupon_total_quantity"`
	DescTxt              string   `json:"desc_txt"`
	GoodsDesc            string   `json:"goods_desc"`
	GoodsGalleryUrls     []string `json:"goods_gallery_urls"`
	GoodsId              int64    `json:"goods_id"`
	GoodsImageUrl        string   `json:"goods_image_url"`
	GoodsName            string   `json:"goods_name"`
	GoodsSign            string   `json:"goods_sign"`
	GoodsThumbnailUrl    string   `json:"goods_thumbnail_url"`
	LgstTxt              string   `json:"lgst_txt"`
	MallId               int64    `json:"mall_id"`
	MallName             string   `json:"mall_name"`
	MerchantType         string   `json:"merchant_type"`
	MinGroupPrice        int64    `json:"min_group_price"`
	MinNormalPrice       int64    `json:"min_normal_price"`
	OptId                int64    `json:"opt_id"`
	OptIds               []int64  `json:"opt_ids"`
	OptName              string   `json:"opt_name"`
	PredictPromotionRate int64    `json:"predict_promotion_rate"`
	PromotionRate        int64    `json:"promotion_rate"`
	SalesTip             string   `json:"sales_tip"`
	SearchId             string   `json:"search_id"`
	ServTxt              string   `json:"serv_txt"`
	ShareRate            int64    `json:"share_rate"`
	CategoryId           int64    `json:"category_id"`
	CategoryName         string   `json:"category_name"`
}
type TopGoodsListQuery struct {
	TopGoodsListGetResponse struct {
		List     []TopGoodsListQueryList `json:"list"`
		ListId   string                  `json:"list_id"`
		SearchId string                  `json:"search_id"`
		Total    int64                   `json:"total"`
	} `json:"top_goods_list_get_response"`
}

// TopGoodsListQuery 多多客获取爆款排行商品接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.top.goods.list.query
func (c *Client) TopGoodsListQuery(ctx context.Context, notMustParams ...*gorequest.Params) (response TopGoodsListQuery, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.top.goods.list.query", notMustParams...)
	params.Set("p_id", c.GetPid())

	// 请求
	err = c.request(ctx, params, &response)
	return
}
