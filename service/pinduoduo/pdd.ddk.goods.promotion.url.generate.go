package pinduoduo

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type GoodsPromotionUrlGenerateResponse struct {
	GoodsPromotionUrlGenerateResponse struct {
		GoodsPromotionUrlList []struct {
			MobileShortUrl string `json:"mobile_short_url,omitempty"` // 对应出参mobile_url的短链接，与mobile_url功能一致。
			MobileUrl      string `json:"mobile_url,omitempty"`       // 使用此推广链接，用户安装微信的情况下，默认拉起拼多多福利券微信小程序，否则唤起H5页面
			QqAppInfo      struct {
				AppId             string `json:"app_id,omitempty"`              // 拼多多小程序id
				BannerUrl         string `json:"banner_url,omitempty"`          // Banner图
				Desc              string `json:"desc,omitempty"`                // 描述
				PagePath          string `json:"page_path,omitempty"`           // 小程序path值
				QqAppIconUrl      string `json:"qq_app_icon_url,omitempty"`     // 小程序icon
				SourceDisplayName string `json:"source_display_name,omitempty"` // 来源名
				Title             string `json:"title,omitempty"`               // 小程序标题
				UserName          string `json:"user_name,omitempty"`           // 用户名
			} `json:"qq_app_info"`
			SchemaUrl string `json:"schema_url,omitempty"` // 使用此推广链接，用户安装拼多多APP的情况下会唤起APP（需客户端支持schema跳转协议）
			ShortUrl  string `json:"short_url,omitempty"`  // 对应出参url的短链接，与url功能一致
			Url       string `json:"url,omitempty"`        // 普通推广长链接，唤起H5页面
			WeAppInfo struct {
				AppId             string `json:"app_id,omitempty"`              // 小程序id
				BannerUrl         string `json:"banner_url,omitempty"`          // Banner图
				Desc              string `json:"desc,omitempty"`                // 描述
				PagePath          string `json:"page_path,omitempty"`           // 小程序path值
				SourceDisplayName string `json:"source_display_name,omitempty"` // 来源名
				Title             string `json:"title,omitempty"`               // 小程序标题
				UserName          string `json:"user_name,omitempty"`           // 用户名
				WeAppIconUrl      string `json:"we_app_icon_url,omitempty"`     // 小程序图片
			} `json:"we_app_info"`
			WeixinCode           string `json:"weixin_code"`
			WeAppWebViewUrl      string `json:"we_app_web_view_url"`
			WeAppWebViewShortUrl string `json:"we_app_web_view_short_url"`
			TzSchemaUrl          string `json:"tz_schema_url"`
			WeixinShortLink      string `json:"weixin_short_link"`
		} `json:"goods_promotion_url_list"`
	} `json:"goods_promotion_url_generate_response"`
}

type GoodsPromotionUrlGenerateResult struct {
	Result GoodsPromotionUrlGenerateResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newGoodsPromotionUrlGenerateResult(result GoodsPromotionUrlGenerateResponse, body []byte, http gorequest.Response) *GoodsPromotionUrlGenerateResult {
	return &GoodsPromotionUrlGenerateResult{Result: result, Body: body, Http: http}
}

// GoodsPromotionUrlGenerate 多多进宝推广链接生成
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.promotion.url.generate
func (c *Client) GoodsPromotionUrlGenerate(ctx context.Context, notMustParams ...*gorequest.Params) (*GoodsPromotionUrlGenerateResult, error) {
	// 参数
	params := NewParamsWithType("pdd.ddk.goods.promotion.url.generate", notMustParams...)
	params.Set("p_id", c.GetPid())
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newGoodsPromotionUrlGenerateResult(GoodsPromotionUrlGenerateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response GoodsPromotionUrlGenerateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newGoodsPromotionUrlGenerateResult(response, request.ResponseBody, request), err
}
