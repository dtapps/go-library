package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type SnsJsCode2sessionResponse struct {
	OpenId     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	Unionid    string `json:"unionid"`     // 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回
	Errcode    string `json:"errcode"`     // 错误码
	Errmsg     string `json:"errmsg"`      // 错误信息
}

type SnsJsCode2sessionResult struct {
	Result SnsJsCode2sessionResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newSnsJsCode2sessionResult(result SnsJsCode2sessionResponse, body []byte, http gorequest.Response) *SnsJsCode2sessionResult {
	return &SnsJsCode2sessionResult{Result: result, Body: body, Http: http}
}

func (c *Client) SnsJsCode2session(ctx context.Context, jsCode string, notMustParams ...gorequest.Params) (*SnsJsCode2sessionResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", c.GetAppId(), c.GetAppSecret(), jsCode), params, http.MethodGet)
	if err != nil {
		return newSnsJsCode2sessionResult(SnsJsCode2sessionResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response SnsJsCode2sessionResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newSnsJsCode2sessionResult(response, request.ResponseBody, request), err
}
