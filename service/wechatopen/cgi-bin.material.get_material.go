package wechatopen

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
func (c *Client) CgiBinMaterialGetMaterial(ctx context.Context, mediaId string) (*CgiBinMaterialGetMaterialResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParams()
	params["media_id"] = mediaId // 要获取的素材的media_id
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/material/get_material?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response CgiBinMaterialGetMaterialResponse
	// 判断内容是否为图片
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		// 可能是图片
		return newCgiBinMaterialGetMaterialResult(CgiBinMaterialGetMaterialResponse{}, request.ResponseBody, request), nil
	}
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
