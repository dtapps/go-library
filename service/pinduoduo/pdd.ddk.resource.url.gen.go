package pinduoduo

import "encoding/json"

type ResourceUrlGenResponse struct {
	ResourceUrlResponse struct {
		MultiUrlList struct {
			ShortUrl string `json:"short_url"` // 频道推广短链接
			Url      string `json:"url"`       // 频道推广长链接
		} `json:"multi_url_list"` // 多人团链接
		Sign          string `json:"sign,omitempty"` // sign
		SingleUrlList struct {
			ShortUrl string `json:"short_url"` // 频道推广短链接
			Url      string `json:"url"`       // 频道推广长链接
		} `json:"single_url_list"` // 单人团链接
		WeAppInfo struct {
			AppId             string `json:"app_id"`              // 小程序id
			BannerUrl         string `json:"banner_url"`          // Banner图
			Desc              string `json:"desc"`                // 描述
			PagePath          string `json:"page_path"`           // 小程序path值
			SourceDisplayName string `json:"source_display_name"` // 来源名
			Title             string `json:"title"`               // 小程序标题
			UserName          string `json:"user_name"`           // 用户名
			WeAppIconUrl      string `json:"we_app_icon_url"`     // 小程序icon
		} `json:"we_app_info"` // 拼多多福利券微信小程序信息
	} `json:"resource_url_response"`
}

type ResourceUrlGenResult struct {
	Result ResourceUrlGenResponse // 结果
	Body   []byte                 // 内容
	Err    error                  // 错误
}

func NewResourceUrlGenResult(result ResourceUrlGenResponse, body []byte, err error) *ResourceUrlGenResult {
	return &ResourceUrlGenResult{Result: result, Body: body, Err: err}
}

// ResourceUrlGen 创建多多进宝推广位
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.pid.generate
func (app *App) ResourceUrlGen(notMustParams ...Params) *ResourceUrlGenResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.resource.url.gen", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var response ResourceUrlGenResponse
	err = json.Unmarshal(body, &response)
	return NewResourceUrlGenResult(response, body, err)
}
