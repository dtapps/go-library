package wechatunion

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PromotionUpdResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type PromotionUpdResult struct {
	Result PromotionUpdResponse // 结果
	Body   []byte               // 内容
	Err    error                // 错误
}

func NewPromotionUpdResult(result PromotionUpdResponse, body []byte, err error) *PromotionUpdResult {
	return &PromotionUpdResult{Result: result, Body: body, Err: err}
}

// PromotionUpd 编辑推广位
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_3-%E7%BC%96%E8%BE%91%E6%8E%A8%E5%B9%BF%E4%BD%8D
func (app *App) PromotionUpd(notMustParams ...Params) *PromotionUpdResult {
	app.AccessToken = app.GetAccessToken()
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request(UnionUrl+fmt.Sprintf("/promoter/promotion/upd?access_token%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response PromotionUpdResponse
	err = json.Unmarshal(body, &response)
	return NewPromotionUpdResult(response, body, err)
}
