package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// CgiBinMaterialGetMaterial 获取永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (c *Client) CgiBinMaterialGetMaterial(ctx context.Context, authorizerAccessToken, mediaId string, notMustParams ...*gorequest.Params) (response APIResponse, body []byte, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("media_id", mediaId) // 要获取的素材的media_id

	// 请求
	body, err = c.requestImage(ctx, "cgi-bin/material/get_material?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
