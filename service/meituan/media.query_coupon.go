package meituan

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type MediaQueryCouponResponse struct {
	Code    int64  `json:"code"`    // 响应码，0成功，其他值为失败
	Message string `json:"message"` // 响应文案
	Data    []struct {
		AvailablePoiInfo struct {
			AvailablePoiNum     int64       `json:"availablePoiNum"` // 可用门店数量。针对到店、到家医药业务类型商品，若传入经纬度信息，则为经纬度所在城市可用的门店数。若不传入经纬度信息，则输出北京可用的门店数
			AvailablePoiCityNum interface{} `json:"availablePoiCityNu,omitempty"`
		} `json:"availablePoiInfo"` // 可用门店信息
		BrandInfo struct {
			BrandName    string `json:"brandName"`    // 品牌名称
			BrandLogoUrl string `json:"brandLogoUrl"` // 品牌Logo的url
		} `json:"brandInfo"` // 品牌信息
		CommissionInfo struct {
			CommissionPercent string `json:"commissionPercent"`    // 查询当时生效的佣金比例， 商品券拉取、通过商品券ID查询、通过榜单listTopiId查询，返回的数据需要除以100表示对应的佣金比例，如返回400表示佣金比例为4%
			Commission        string `json:"commission,omitempty"` // 只支持到店、到家医药业务类型。查询当时生效的佣金值。单位元，保留小数点后两位
		} `json:"commissionInfo"` // 佣金信息
		CouponPackDetail struct {
			Name          string `json:"name"`                // 商品名称
			SkuViewId     string `json:"skuViewId"`           // 商品skuViewId，传入开放平台取链接口的skuViewId，取得对应推广链接才能正常归因订单
			Specification string `json:"specification"`       // 规格信息，只支持到家医药商品业务类型
			CouponNum     int64  `json:"couponNum,omitempty"` // 只支持到家外卖商品券业务类型，券包中券的数量
			ValidTime     int64  `json:"validTime,omitempty"` // 只支持到家外卖商品券业务类型，活动截止有效日期，仅作参考，具体结束时间详见couponValidTimeInfo中的信息
			HeadUrl       string `json:"headUrl"`             // 商品头图的url
			SaleVolume    string `json:"saleVolume"`          // 美团累计销量，例：100+，1000+，10000+
			StartTime     int64  `json:"startTime,omitempty"` // 只支持到家外卖商品券业务类型，活动有效期开始时间
			EndTime       int64  `json:"endTime,omitempty"`   // 只支持到家外卖商品券业务类型，活动有效期结束时间
			SaleStatus    bool   `json:"saleStatus"`          // 售卖状态，可售为是，不可售为否。不可售商品不返回商品数据
			OriginalPrice string `json:"originalPrice"`       // 原始价格，如划线价(元）
			SellPrice     string `json:"sellPrice"`           // 售卖价格(元）
			Platform      int64  `json:"platform"`            // 平台，1-到家、2-到店
			BizLine       int64  `json:"bizLine"`             // 二级分类，当platform为1时null代表外卖，当platform为2时1代表餐
			RecallSource  any    `json:"recallSource,omitempty"`
		} `json:"couponPackDetail"` // 商品详情
		DeliverablePoiInfo struct {
			PoiName          string `json:"poiName"`          // 门店名称，商品券可配送门店信息，无则不返回 注：入参经纬度可展示附近配送门店名称。按主题榜单查询时不展示该字段
			PoiLogoUrl       string `json:"poiLogoUrl"`       // 门店Logo的url 注：入参经纬度可展示附近配送门店logo。按主题榜单查询时不展示该字段。
			DeliveryDistance string `json:"deliveryDistance"` // 配送距离 注：入参经纬度可展示附近配送门店的配送距离。按主题榜单查询时不展示该字段。
			DistributionCost string `json:"distributionCost"` // 配送费 注：入参经纬度可展示附近配送门店的配送费。按主题榜单查询时不展示该字段。
			DeliveryDuration string `json:"deliveryDuration"` // 配送时长 注：入参经纬度可展示附近配送门店的配送时长。按主题榜单查询时不展示该字段。
			LastDeliveryFee  string `json:"lastDeliveryFee"`  // 起送额 注：入参经纬度可展示附近配送门店的起送金额。按主题榜单查询时不展示该字段。
		} `json:"deliverablePoiInfo,omitempty"` // 只支持到家外卖商品券业务类型，可配送门店信息
		PurchaseLimitInfo struct {
			SingleDayPurchaseLimit int64 `json:"singleDayPurchaseLimit"` // 单日售卖上限
		} `json:"purchaseLimitInfo"` // 	购买限制信息
		CouponValidTimeInfo struct {
			CouponValidTimeType int64 `json:"couponValidTimeType"` // 券包活动生效时间类型,1:按生效天数,2:按时间段
			CouponValidDay      int64 `json:"couponValidDay"`      // 券生效天数；couponValidTimeType为1有效
			CouponValidSTime    int64 `json:"couponValidSTime"`    // 券开始时间戳，单位秒；couponValidTimeType为2有效
			CouponValidETime    int64 `json:"couponValidETime"`    // 券结束时间戳，单位秒；couponValidTimeType为2有效
		} `json:"couponValidTimeInfo,omitempty"` // 只支持到家外卖商品券业务类型，券包活动有效时间信息
	} `json:"data,omitempty"` // 响应结果信息
	HasNext  bool   `json:"hasNext,omitempty"`  // 分页使用，看是否有下一页
	SearchId string `json:"searchId,omitempty"` // 搜索场景出参,用于相同条件下一页请求入参
}

type MediaQueryCouponResult struct {
	Result MediaQueryCouponResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newMediaQueryCouponResult(result MediaQueryCouponResponse, body []byte, http gorequest.Response) *MediaQueryCouponResult {
	return &MediaQueryCouponResult{Result: result, Body: body, Http: http}
}

// MediaQueryCoupon 商品查询接口
// 查询售卖商品接口，支持全量查询、精确查询、榜单主题查询。需用POST方式调用。只接受JSON格式。
// https://media.meituan.com/pc/index.html#/materials/api-detail/query_coupon
func (c *MediaClient) MediaQueryCoupon(ctx context.Context, notMustParams ...*gorequest.Params) (*MediaQueryCouponResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response MediaQueryCouponResponse
	request, err := c.request(ctx, "cps_open/common/api/v1/query_coupon", http.MethodPost, params, &response)
	return newMediaQueryCouponResult(response, request.ResponseBody, request), err
}
