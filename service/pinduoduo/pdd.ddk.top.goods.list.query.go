package pinduoduo

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type TopGoodsListQueryResponse struct {
	TopGoodsListGetResponse struct {
		List []struct {
			CatIds               []int    `json:"cat_ids"`
			CouponDiscount       int      `json:"coupon_discount"`
			CouponEndTime        int      `json:"coupon_end_time"`
			CouponMinOrderAmount int      `json:"coupon_min_order_amount"`
			CouponRemainQuantity int      `json:"coupon_remain_quantity"`
			CouponStartTime      int      `json:"coupon_start_time"`
			CouponTotalQuantity  int      `json:"coupon_total_quantity"`
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
			MinGroupPrice        int      `json:"min_group_price"`
			MinNormalPrice       int      `json:"min_normal_price"`
			OptId                int      `json:"opt_id"`
			OptIds               []int    `json:"opt_ids"`
			OptName              string   `json:"opt_name"`
			PredictPromotionRate int      `json:"predict_promotion_rate"`
			PromotionRate        int      `json:"promotion_rate"`
			SalesTip             string   `json:"sales_tip"`
			SearchId             string   `json:"search_id"`
			ServTxt              string   `json:"serv_txt"`
			ShareRate            int      `json:"share_rate"`
			CategoryId           int64    `json:"category_id"`
			CategoryName         string   `json:"category_name"`
		} `json:"list"`
		ListId   string `json:"list_id"`
		SearchId string `json:"search_id"`
		Total    int64  `json:"total"`
	} `json:"top_goods_list_get_response"`
}
type TopGoodsListQueryResult struct {
	Result TopGoodsListQueryResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func newTopGoodsListQueryResult(result TopGoodsListQueryResponse, body []byte, http gorequest.Response, err error) *TopGoodsListQueryResult {
	return &TopGoodsListQueryResult{Result: result, Body: body, Http: http, Err: err}
}

// TopGoodsListQuery 多多客获取爆款排行商品接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.top.goods.list.query
func (c *Client) TopGoodsListQuery(ctx context.Context, notMustParams ...Params) *TopGoodsListQueryResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.top.goods.list.query", notMustParams...)
	params.Set("p_id", c.GetPid())
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var response TopGoodsListQueryResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTopGoodsListQueryResult(response, request.ResponseBody, request, err)
}
