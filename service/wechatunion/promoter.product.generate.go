package wechatunion

import (
	"fmt"
	"net/http"
)

type PromoterProductGenerateResult struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	List    []struct {
		ProductId   string `json:"productId"` // 商品SPU ID
		Pid         string `json:"pid"`       // 推广位PID
		ProductInfo struct {
			ProductId     string   `json:"productId"`     // 商品SPU ID
			Title         string   `json:"title"`         // 商品标题
			SubTitle      string   `json:"subTitle"`      // 商品子标题
			HeadImg       []string `json:"headImg"`       // 商品主图
			MinPrice      int      `json:"minPrice"`      // 商品最低价格，单位分
			Discount      int      `json:"discount"`      // 商品优惠金额，单位分
			DiscountPrice int      `json:"discountPrice"` // 商品券后最低价格，单位分
			ShopName      string   `json:"shopName"`      // 商店名称
			PluginResult  int      `json:"pluginResult"`  // 是否引用小商店组件（未引用组件的商品不可推广），0：否，1：是
			TotalStockNum int      `json:"totalStockNum"` // 商品库存
		} `json:"productInfo"` // 商品相关信息
		ShareInfo struct {
			AppId                  string `json:"appId"`                  // 推广商品的小程序AppID
			Path                   string `json:"path"`                   // 推广商品的小程序Path
			CouponPath             string `json:"couponPath"`             // 推广商品的带券小程序Path
			WxaCode                string `json:"wxaCode"`                // 已废弃。推广商品详情页的不带券葵花码图片
			CouponWxaCode          string `json:"couponWxaCode"`          // 已废弃。推广商品详情页的带券葵花码图片
			PromotionUrl           string `json:"promotionUrl"`           // 推广商品短链
			CouponPromotionUrl     string `json:"couponPromotionUrl"`     // 推广商品带券短链
			PromotionWording       string `json:"promotionWording"`       // 推广商品文案
			CouponPromotionWording string `json:"couponPromotionWording"` // 推广商品带券文案
		} `json:"shareInfo"` // 推广相关信息
	} `json:"list"`
}

// PromoterProductGenerate
// 获取商品推广素材
// 通过该接口获取商品的推广素材，包括店铺appID、商品详情页Path、推广文案及推广短链、商品图片等
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/product/category.html
func (app *App) PromoterProductGenerate(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err = app.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/product/generate?access_token=%s", app.AccessToken), params, http.MethodPost)
	return body, err
}
