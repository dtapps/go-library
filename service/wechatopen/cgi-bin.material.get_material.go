package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinMaterialGetMaterialResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CgiBinMaterialGetMaterialResult struct {
	Result CgiBinMaterialGetMaterialResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newCgiBinMaterialGetMaterialResult(result CgiBinMaterialGetMaterialResponse, body []byte, http gorequest.Response) *CgiBinMaterialGetMaterialResult {
	return &CgiBinMaterialGetMaterialResult{Result: result, Body: body, Http: http}
}

// CgiBinMaterialGetMaterial 获取永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (c *Client) CgiBinMaterialGetMaterial(ctx context.Context, authorizerAccessToken, mediaId string, notMustParams ...*gorequest.Params) (*CgiBinMaterialGetMaterialResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("media_id", mediaId) // 要获取的素材的media_id

	// 请求
	var response CgiBinMaterialGetMaterialResponse
	request, err := c.request(ctx, "cgi-bin/material/get_material?access_token="+authorizerAccessToken, params, http.MethodPost, &response)

	// 判断内容是否为图片
	//if request.HeaderIsImg() == false {
	//	err = json.Unmarshal(request.ResponseBody, &response)
	//}
	return newCgiBinMaterialGetMaterialResult(response, request.ResponseBody, request), err
}
