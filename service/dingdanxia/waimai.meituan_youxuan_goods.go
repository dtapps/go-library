package dingdanxia

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WaiMaiMeituanYouxuanGoodsResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Total   int64 `json:"total"` // 数据总量
		skuList []struct {
			ItemId                string   `json:"itemId"`                // 商品信息-宝贝 id skuid
			Title                 string   `json:"title"`                 // 商品信息-商品标题
			PictUrl               string   `json:"pictUrl"`               // 商品信息-商品主图
			SmallImages           []string `json:"smallImages"`           // 商品信息-商品小图列表
			OriginPrice           string   `json:"originPrice"`           // 原价（单位为元）
			PromotionPrice        string   `json:"promotionPrice"`        // 促销价（活动价格）秒杀价格 （单位为元）
			ItemDeepLinkUrl       string   `json:"itemDeepLinkUrl"`       // 宝贝的 deeplink 地址
			ItemMiddlePageLinkUrl string   `json:"itemMiddlePageLinkUrl"` // 宝贝的中间页地址
			ItemWXLinkUrl         string   `json:"itemWXLinkUrl"`         // 宝贝的微信小程序链接地址
			HotFlag               bool     `json:"hotFlag"`               // 是否是热门商品
		}
	} `json:"data"`
}

type WaiMaiMeituanYouxuanGoodsResult struct {
	Result WaiMaiMeituanYouxuanGoodsResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newWaiMaiMeituanYouxuanGoodsResult(result WaiMaiMeituanYouxuanGoodsResponse, body []byte, http gorequest.Response) *WaiMaiMeituanYouxuanGoodsResult {
	return &WaiMaiMeituanYouxuanGoodsResult{Result: result, Body: body, Http: http}
}

// WaiMaiMeituanYouxuanGoods 优选商品查询API【2022年1月17日暂停数据访问】
// https://www.dingdanxia.com/doc/235/173
func (c *Client) WaiMaiMeituanYouxuanGoods(ctx context.Context, notMustParams ...*gorequest.Params) (*WaiMaiMeituanYouxuanGoodsResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/waimai/meituan_youxuan_goods", params, http.MethodPost)
	if err != nil {
		return newWaiMaiMeituanYouxuanGoodsResult(WaiMaiMeituanYouxuanGoodsResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WaiMaiMeituanYouxuanGoodsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWaiMaiMeituanYouxuanGoodsResult(response, request.ResponseBody, request), err
}
