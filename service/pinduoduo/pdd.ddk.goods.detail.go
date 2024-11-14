package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
)

type GoodsDetailGoodsDetailsResponse struct {
	ActivityPromotionRate       int64    `json:"activity_promotion_rate,omitempty"`         // 活动佣金比例，千分比（特定活动期间的佣金比例）
	ActivityTags                []int64  `json:"activity_tags,omitempty"`                   // 商品活动标记数组，例：[4,7]，4-秒杀 7-百亿补贴等
	BrandName                   string   `json:"brand_name,omitempty"`                      // 商品品牌词信息，如“苹果”、“阿迪达斯”、“李宁”等
	CashGiftAmount              int64    `json:"cash_gift_amount,omitempty"`                // 全局礼金金额，单位分
	CatId                       int64    `json:"cat_id,omitempty"`                          // 商品类目ID，使用pdd.goods.cats.get接口获取
	CatIds                      []int64  `json:"cat_ids,omitempty"`                         // 商品一~四级类目ID列表
	CltCpnBatchSn               string   `json:"clt_cpn_batch_sn,omitempty"`                // 店铺收藏券id
	CltCpnDiscount              int64    `json:"clt_cpn_discount,omitempty"`                // 店铺收藏券面额,单位为分
	CltCpnEndTime               int64    `json:"clt_cpn_end_time,omitempty"`                // 店铺收藏券截止时间
	CltCpnMinAmt                int64    `json:"clt_cpn_min_amt,omitempty"`                 // 店铺收藏券使用门槛价格,单位为分
	CltCpnQuantity              int64    `json:"clt_cpn_quantity,omitempty"`                // 店铺收藏券总量
	CltCpnRemainQuantity        int64    `json:"clt_cpn_remain_quantity,omitempty"`         // 店铺收藏券剩余量
	CltCpnStartTime             int64    `json:"clt_cpn_start_time,omitempty"`              // 店铺收藏券起始时间
	CouponDiscount              int64    `json:"coupon_discount,omitempty"`                 // 优惠券面额，单位为分
	CouponEndTime               int64    `json:"coupon_end_time,omitempty"`                 // 优惠券失效时间，UNIX时间戳
	CouponMinOrderAmount        int64    `json:"coupon_min_order_amount,omitempty"`         // 优惠券门槛金额，单位为分
	CouponRemainQuantity        int64    `json:"coupon_remain_quantity,omitempty"`          // 优惠券剩余数量
	CouponStartTime             int64    `json:"coupon_start_time,omitempty"`               // 优惠券生效时间，UNIX时间戳
	CouponTotalQuantity         int64    `json:"coupon_total_quantity,omitempty"`           // 优惠券总数量
	CreateAt                    int64    `json:"create_at,omitempty"`                       // 创建时间（unix时间戳）
	DescTxt                     string   `json:"desc_txt,omitempty"`                        // 描述分
	ExtraCouponAmount           int64    `json:"extra_coupon_amount,omitempty"`             // 额外优惠券
	GoodsDesc                   string   `json:"goods_desc,omitempty"`                      // 参与多多进宝的商品描述
	GoodsGalleryUrls            []string `json:"goods_gallery_urls,omitempty"`              // 商品轮播图
	GoodsImageUrl               string   `json:"goods_image_url,omitempty"`                 // 多多进宝商品主图
	GoodsName                   string   `json:"goods_name,omitempty"`                      // 参与多多进宝的商品标题
	GoodsSign                   string   `json:"goods_sign,omitempty"`                      // 商品goodsSign，支持通过goodsSign查询商品。goodsSign是加密后的goodsId, goodsId已下线，请使用goodsSign来替代。使用说明：https://jinbao.pinduoduo.com/qa-system?questionId=252
	GoodsThumbnailUrl           string   `json:"goods_thumbnail_url,omitempty"`             // 商品缩略图
	HasCoupon                   bool     `json:"has_coupon,omitempty"`                      // 商品是否有优惠券 true-有，false-没有
	HasMallCoupon               bool     `json:"has_mall_coupon,omitempty"`                 // 是否有店铺券
	LgstTxt                     string   `json:"lgst_txt,omitempty"`                        // 物流分
	MallCouponDiscountPct       int64    `json:"mall_coupon_discount_pct,omitempty"`        // 店铺折扣
	MallCouponEndTime           int64    `json:"mall_coupon_end_time,omitempty"`            // 店铺券使用结束时间
	MallCouponMaxDiscountAmount int64    `json:"mall_coupon_max_discount_amount,omitempty"` // 最大使用金额
	MallCouponMinOrderAmount    int64    `json:"mall_coupon_min_order_amount,omitempty"`    // 最小使用金额
	MallCouponRemainQuantity    int64    `json:"mall_coupon_remain_quantity,omitempty"`     // 店铺券余量
	MallCouponStartTime         int64    `json:"mall_coupon_start_time,omitempty"`          // 店铺券使用开始时间
	MallCouponTotalQuantity     int64    `json:"mall_coupon_total_quantity,omitempty"`      // 店铺券总量
	MallCps                     int64    `json:"mall_cps,omitempty"`                        // 该商品所在店铺是否参与全店推广，0：否，1：是
	MallId                      int64    `json:"mall_id,omitempty"`                         // 商家id
	MallImgUrl                  string   `json:"mall_img_url,omitempty"`                    // 店铺logo图
	MallName                    string   `json:"mall_name,omitempty"`                       // 店铺名称
	MaterialList                []struct {
		Id           string   `json:"id,omitempty"`            // 素材ID
		ImageList    []string `json:"image_list,omitempty"`    // 图片列表
		TextList     []string `json:"text_list,omitempty"`     // 文字列表
		ThumbnailUrl string   `json:"thumbnail_url,omitempty"` // 视频缩略图
		Mtype        int64    `json:"type,omitempty"`          // 素材类型，1-图文，2-视频
		VideoUrl     string   `json:"video_url,omitempty"`     // 视频url
	} `json:"material_list"`
	MerchantType               int64    `json:"merchant_type,omitempty"`                  // 店铺类型，1-个人，2-企业，3-旗舰店，4-专卖店，5-专营店，6-普通店（未传为全部）
	MinGroupPrice              int64    `json:"min_group_price,omitempty"`                // 最低价sku的拼团价，单位为分
	MinNormalPrice             int64    `json:"min_normal_price,omitempty"`               // 最低价sku的单买价，单位为分
	OnlySceneAuth              bool     `json:"only_scene_auth,omitempty"`                // 快手专享
	OptId                      int64    `json:"opt_id,omitempty"`                         // 商品标签ID，使用pdd.goods.opt.get接口获取
	OptIds                     []int64  `json:"opt_ids,omitempty"`                        // 商品标签ID
	OptName                    string   `json:"opt_name,omitempty"`                       // 商品标签名称
	PlanType                   int64    `json:"plan_type,omitempty"`                      // 推广计划类型: 1-全店推广 2-单品推广 3-定向推广 4-招商推广 5-分销推广
	PredictPromotionRate       int64    `json:"predict_promotion_rate,omitempty"`         // 比价行为预判定佣金，需要用户备案
	PromotionRate              int64    `json:"promotion_rate,omitempty"`                 // 佣金比例，千分比
	SalesTip                   string   `json:"sales_tip,omitempty"`                      // 已售卖件数
	ServiceTags                []int64  `json:"service_tags,omitempty"`                   // 服务标签: 4-送货入户并安装,5-送货入户,6-电子发票,9-坏果包赔,11-闪电退款,12-24小时发货,13-48小时发货,17-顺丰包邮,18-只换不修,1可定制化,29-预约配送,1000001-正品发票,1000002-送货入户并安装
	ServTxt                    string   `json:"serv_txt,omitempty"`                       // 服务分
	ShareRate                  int64    `json:"share_rate,omitempty"`                     // 招商分成服务费比例，千分比
	SubsidyAmount              int64    `json:"subsidy_amount,omitempty"`                 // 优势渠道专属商品补贴金额，单位为分。针对优质渠道的补贴活动，指定优势渠道可通过推广该商品获取相应补贴。补贴活动入口：[进宝网站-官方活动-千万补贴]，报名入口：https://jinbao.pinduoduo.com/ten-million-subsidy/entry
	SubsidyDuoAmountTenMillion int64    `json:"subsidy_duo_amount_ten_million,omitempty"` // 千万补贴给渠道的收入补贴，不允许直接给下级代理展示，单位为分
	UnifiedTags                []string `json:"unified_tags,omitempty"`                   // 优惠标签列表，包括："X元券","比全网低X元","服务费","精选素材","近30天低价","同款低价","同款好评","同款热销","旗舰店","一降到底","招商优选","商家优选","好价再降X元","全站销量XX","实时热销榜第X名","实时好评榜第X名","额外补X元"等
	VideoUrls                  []string `json:"video_urls,omitempty"`                     // 商品视频url
	ZsDuoId                    int64    `json:"zs_duo_id,omitempty"`                      // 招商团长id
	GoodsId                    int64    `json:"goods_id"`
	CategoryId                 int64    `json:"category_id"`
	CategoryName               string   `json:"category_name"`
}

type GoodsDetailResponse struct {
	GoodsDetailResponse struct {
		GoodsDetails []GoodsDetailGoodsDetailsResponse `json:"goods_details"`
	} `json:"goods_detail_response"`
}
type GoodsDetailResponseError struct {
	ErrorResponse struct {
		SubMsg    string `json:"sub_msg"`
		SubCode   string `json:"sub_code"`
		ErrorMsg  string `json:"error_msg"`
		ErrorCode int    `json:"error_code"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}

type GoodsDetailResult struct {
	Result GoodsDetailResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newGoodsDetailResult(result GoodsDetailResponse, body []byte, http gorequest.Response) *GoodsDetailResult {
	return &GoodsDetailResult{Result: result, Body: body, Http: http}
}

// GoodsDetail 多多进宝商品详情查询
// https://open.pinduoduo.com/application/document/api?id=pdd.ddk.goods.detail
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.detail
func (c *Client) GoodsDetail(ctx context.Context, notMustParams ...*gorequest.Params) (*GoodsDetailResult, GoodsDetailResponseError, error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.goods.detail", notMustParams...)
	params.Set("pid", c.GetPid())

	// 请求
	var response GoodsDetailResponse
	request, err := c.request(ctx, params, &response)
	var responseError GoodsDetailResponseError
	_ = gojson.Unmarshal(request.ResponseBody, &responseError)
	return newGoodsDetailResult(response, request.ResponseBody, request), responseError, err
}
