package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetSettingCategoriesResponse struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	Categories []struct {
		First       int    `json:"first"`        // 一级类目 ID
		FirstName   string `json:"first_name"`   // 一级类目名称
		Second      int    `json:"second"`       // 二级类目 ID
		SecondName  string `json:"second_name"`  // 二级类目名称
		AuditStatus int    `json:"audit_status"` // 审核状态（1 审核中 2 审核不通过 3 审核通过）
		AuditReason string `json:"audit_reason"` // 审核不通过的原因
	} `json:"categories"` // 已设置的类目信息列表
	Limit         int `json:"limit"`          // 一个更改周期内可以添加类目的次数
	Quota         int `json:"quota"`          // 本更改周期内还可以添加类目的次数
	CategoryLimit int `json:"category_limit"` // 最多可以设置的类目数量
}

type GetSettingCategoriesResult struct {
	Result GetSettingCategoriesResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newGetSettingCategoriesResult(result GetSettingCategoriesResponse, body []byte, http gorequest.Response) *GetSettingCategoriesResult {
	return &GetSettingCategoriesResult{Result: result, Body: body, Http: http}
}

// GetSettingCategories 获取已设置的所有类目
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/category-management/getSettingCategories.html
func (c *Client) GetSettingCategories(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetSettingCategoriesResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetSettingCategoriesResponse
	request, err := c.request(ctx, "cgi-bin/wxopen/getcategory?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetSettingCategoriesResult(response, request.ResponseBody, request), err
}
