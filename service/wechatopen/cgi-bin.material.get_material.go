package wechatopen

import (
	"bytes"
	"context"
	"errors"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostorage"
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
func (c *Client) CgiBinMaterialGetMaterial(ctx context.Context, mediaId string, notMustParams ...gorequest.Params) (*CgiBinMaterialGetMaterialResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newCgiBinMaterialGetMaterialResult(CgiBinMaterialGetMaterialResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("media_id", mediaId) // 要获取的素材的media_id
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/material/get_material?access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newCgiBinMaterialGetMaterialResult(CgiBinMaterialGetMaterialResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinMaterialGetMaterialResponse
	// 判断内容是否为图片
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinMaterialGetMaterialResult(response, request.ResponseBody, request), err
}

func (cr *CgiBinMaterialGetMaterialResult) SaveImg(db *gostorage.AliYun, fileName, filePath string) error {
	if cr.Result.Errcode != 0 {
		return errors.New(cr.Result.Errmsg)
	}
	// 上传
	_, err := db.PutObject(bytes.NewReader(cr.Body), filePath, fileName)
	if err != nil {
		return err
	}
	return nil
}
