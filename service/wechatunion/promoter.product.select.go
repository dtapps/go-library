package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PromoterProductSelectResponse struct {
	Errcode     int    `json:"errcode"` // 错误码
	Errmsg      string `json:"errmsg"`  // 错误信息
	Total       int64  `json:"total"`   // 商品总数
	ProductList []struct {
		ProductId string `json:"productId"` // 商品SPU ID
		Product   struct {
			ProductId string `json:"productId"` // 商品SPU ID
			Info      struct {
				Title    string   `json:"title"`    // 商品标题
				SubTitle string   `json:"subTitle"` // 商品子标题
				HeadImg  []string `json:"headImg"`  // 商品主图
				Category []struct {
					CatId string `json:"catId"` // 类目ID
					Name  string `json:"name"`  // 类目ID
				} `json:"category"` // 商品类目
				Brand   string `json:"brand"`   // 品牌名称
				BrandId string `json:"brandId"` // 品牌ID
				Model   string `json:"model"`   // 型号
				Detail  struct {
					DetailImg []string `json:"detailImg"` // 	商品详情图片
				} `json:"detail"` // 商品详细数据
				Param         []interface{} `json:"param"`         // 商品参数
				MinPrice      int64         `json:"minPrice"`      // 商品最低价格，单位分
				TotalStockNum int64         `json:"totalStockNum"` // 总库存
				TotalSoldNum  int           `json:"totalSoldNum"`  // 累计销量
				TotalOrderNum int           `json:"totalOrderNum"` // 累计订单量
				DiscountPrice int64         `json:"discountPrice"` // 商品券后价
			} `json:"info"` // 商品具体信息
			Skus []struct {
				SkuId          string `json:"skuId"` // 商品SKU ID
				ProductSkuInfo struct {
					ThumbImg    string `json:"thumbImg"`              // 商品SKU 小图
					SalePrice   int    `json:"salePrice"`             // 商品SKU 销售价格，单位分
					MarketPrice int    `json:"marketPrice,omitempty"` // 商品SKU 市场价格，单位分
					StockInfo   struct {
						StockNum int `json:"stockNum"` // 	商品SKU 库存
					} `json:"stockInfo"`
				} `json:"productSkuInfo"`
			} `json:"skus"` // 商品SKU
		} `json:"product"` // 商品数据
		LeagueExInfo struct {
			HasCommission   int   `json:"hasCommission"`   // 是否有佣金，1/0
			CommissionRatio int64 `json:"commissionRatio"` // 佣金比例，万分之一
			CommissionValue int64 `json:"commissionValue"` // 佣金金额，单位分
		} `json:"leagueExInfo"` // 联盟佣金相关数据
		ShopInfo struct {
			Name        string `json:"name"`       // 小商店名称
			AppId       string `json:"appId"`      // 小商店AppID
			Username    string `json:"username"`   // 小商店原始id
			HeadImgUrl  string `json:"headImgUrl"` // 小商店店铺头像
			AddressList []struct {
				AddressInfo struct {
					ProvinceName string `json:"provinceName"` // 国标收货地址第一级地址
					CityName     string `json:"cityName"`     // 国标收货地址第二级地址
					CountyName   string `json:"countyName"`   // 国标收货地址第三级地址
				} `json:"addressInfo"` // 地址信息
				AddressType struct {
					Express  int `json:"express"`  // 是否支持快递，1：是，0：否
					SameCity int `json:"sameCity"` // 是否支持同城配送，1：是，0：否
					Pickup   int `json:"pickup"`   // 是否支持上门自提，1：是，0：否
				} `json:"addressType"` // 地址类型
			} `json:"addressList"` // 发货地，只有当配送方式包含「同城配送、上门自提」才出该项
			ShippingMethods struct {
				Express  int `json:"express"`  // 是否支持快递，1：是，0：否
				SameCity int `json:"sameCity"` // 是否支持同城配送，1：是，0：否
				Pickup   int `json:"pickup"`   // 是否支持上门自提，1：是，0：否
			} `json:"shippingMethods"` // 配送方式
			SameCityTemplate struct {
				DeliverScopeType int    `json:"deliverScopeType"` // 配送范围的定义方式，0：按照距离定义配送范围，1：按照区域定义配送范围
				Scope            string `json:"scope"`            // 配送范围
				Region           struct {
					ProvinceName string `json:"provinceName"` // 国标收货地址第一级地址
					CityName     string `json:"cityName"`     // 国标收货地址第二级地址
					CountyName   string `json:"countyName"`   // 国标收货地址第三级地址
				} `json:"region"` // 全城配送时的配送范围
			} `json:"sameCityTemplate"` // 配送范围，只有当配送方式包含「同城配送」才出该项
			FreightTemplate struct {
				NotSendArea struct {
					AddressInfoList []struct {
						ProvinceName string `json:"provinceName"` // 国标收货地址第一级地址
						CityName     string `json:"cityName"`     // 国标收货地址第二级地址
						CountyName   string `json:"countyName"`   // 国标收货地址第三级地址
					} `json:"addressInfoList"` // 不发货地区地址列表
				} `json:"notSendArea,omitempty"` // 不发货地区
			} `json:"freightTemplate"` // 运费模板，只有当配送方式包含「快递」才出此项
		} `json:"shopInfo"` // 商品所属小商店数据
		CouponInfo struct {
			HasCoupon    int    `json:"hasCoupon"` // 是否有联盟券，1为含券商品，0为全部商品
			CouponId     string `json:"couponId"`  // 券id
			CouponDetail struct {
				RestNum      int `json:"restNum"` // 券库存
				Type         int `json:"type"`    // 券类型
				DiscountInfo struct {
					DiscountCondition struct {
						ProductIds   []string `json:"productIds"`   // 指定商品 id
						ProductCnt   string   `json:"productCnt"`   // 商品数
						ProductPrice string   `json:"productPrice"` // 商品金额
					} `json:"discountCondition"` // 指定商品 id
					DiscountNum int   `json:"discountNum,omitempty"` // 折扣数，如 5.1 折 为 5.1 * 1000
					DiscountFee int64 `json:"discountFee,omitempty"` // 直减金额，单位为分
				} `json:"discountInfo"` // 券面额
				ValidInfo struct {
					ValidType   int    `json:"validType"`   // 有效期类型，1 为商品指定时间区间，2 为生效天数
					ValidDayNum int    `json:"validDayNum"` // 生效天数
					StartTime   string `json:"startTime"`   // 有效开始时间
					EndTime     string `json:"endTime"`     // 有效结束时间
				} `json:"validInfo"` // 有效期
				ReceiveInfo struct {
					StartTime         string `json:"startTime"`         // 有效结束时间
					EndTime           string `json:"endTime"`           // 领取结束时间戳
					LimitNumOnePerson int    `json:"limitNumOnePerson"` // 每人限领张数
				} `json:"receiveInfo"` // 领券时间
			} `json:"couponDetail"` // 券详情
		} `json:"couponInfo"` // 联盟优惠券数据
	} `json:"productList"` // 商品列表数据
}

type PromoterProductSelectResult struct {
	Result PromoterProductSelectResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func newPromoterProductSelectResult(result PromoterProductSelectResponse, body []byte, http gorequest.Response, err error) *PromoterProductSelectResult {
	return &PromoterProductSelectResult{Result: result, Body: body, Http: http, Err: err}
}

// PromoterProductSelect
// 查询联盟精选商品
// 支持开发者根据多种筛选条件获取联盟精选的商品列表及详情，筛选条件包括商品价格、商品佣金、商品累计销量、佣金比例、是否含有联盟券、配送方式、发货地区
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/product/category.html#3.%E6%9F%A5%E8%AF%A2%E8%81%94%E7%9B%9F%E7%B2%BE%E9%80%89%E5%95%86%E5%93%81
func (c *Client) PromoterProductSelect(notMustParams ...gorequest.Params) *PromoterProductSelectResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/promoter/product/select?access_token=%s", c.getAccessToken()), params, http.MethodGet)
	// 定义
	var response PromoterProductSelectResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPromoterProductSelectResult(response, request.ResponseBody, request, err)
}
