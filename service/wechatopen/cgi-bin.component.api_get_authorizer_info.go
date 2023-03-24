package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiGetAuthorizerInfoResponse struct {
	AuthorizerInfo struct {
		NickName        string `json:"nick_name"` // 昵称
		HeadImg         string `json:"head_img"`  // 头像
		ServiceTypeInfo struct {
			Id int `json:"id"` // 0=普通小程序 2=门店小程序 3=门店小程序 4=小游戏 10=小商店 12=试用小程序
		} `json:"service_type_info"` // 小程序类型
		VerifyTypeInfo struct {
			Id int `json:"id"` // -1=未认证 0=微信认证
		} `json:"verify_type_info"` // 小程序认证类型
		UserName      string `json:"user_name"`      // 原始 ID
		PrincipalName string `json:"principal_name"` // 主体名称
		Signature     string `json:"signature"`      // 帐号介绍
		BusinessInfo  struct {
			OpenPay   int `json:"open_pay"`
			OpenShake int `json:"open_shake"`
			OpenScan  int `json:"open_scan"`
			OpenCard  int `json:"open_card"`
			OpenStore int `json:"open_store"`
		} `json:"business_info"` // 用以了解功能的开通状况（0代表未开通，1代表已开通)
		QrcodeUrl       string `json:"qrcode_url"` // 二维码图片的 URL，开发者最好自行也进行保存
		MiniProgramInfo struct {
			Network struct {
				RequestDomain      []string      `json:"RequestDomain"`
				WsRequestDomain    []string      `json:"WsRequestDomain"`
				UploadDomain       []string      `json:"UploadDomain"`
				DownloadDomain     []string      `json:"DownloadDomain"`
				BizDomain          []string      `json:"BizDomain"`
				UDPDomain          []string      `json:"UDPDomain"`
				TCPDomain          []interface{} `json:"TCPDomain"`
				NewRequestDomain   []interface{} `json:"NewRequestDomain"`
				NewWsRequestDomain []interface{} `json:"NewWsRequestDomain"`
				NewUploadDomain    []interface{} `json:"NewUploadDomain"`
				NewDownloadDomain  []interface{} `json:"NewDownloadDomain"`
				NewBizDomain       []interface{} `json:"NewBizDomain"`
				NewUDPDomain       []interface{} `json:"NewUDPDomain"`
				NewTCPDomain       []interface{} `json:"NewTCPDomain"`
			} `json:"network"` // 小程序配置的合法域名信息
			Categories []struct {
				First  string `json:"first"`
				Second string `json:"second"`
			} `json:"categories"` // 小程序配置的类目信息
			VisitStatus int `json:"visit_status"`
		} `json:"MiniProgramInfo"` // 小程序配置，根据这个字段判断是否为小程序类型授权
		Alias string `json:"alias"` // 公众号所设置的微信号，可能为空
		Idc   int    `json:"idc"`
	} `json:"authorizer_info"` // 小程序帐号信息
	AuthorizationInfo struct {
		AuthorizerAppid string `json:"authorizer_appid"` // 授权方 appid
		FuncInfo        []struct {
			FuncscopeCategory struct {
				Id int `json:"id"`
			} `json:"funcscope_category"`
			ConfirmInfo struct {
				NeedConfirm    int `json:"need_confirm"`
				AlreadyConfirm int `json:"already_confirm"`
				CanConfirm     int `json:"can_confirm"`
			} `json:"confirm_info,omitempty"`
		} `json:"func_info"` // 授权给开发者的权限集列表
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	} `json:"authorization_info"` // 授权信息
}

type CgiBinComponentApiGetAuthorizerInfoResult struct {
	Result CgiBinComponentApiGetAuthorizerInfoResponse // 结果
	Body   []byte                                      // 内容
	Http   gorequest.Response                          // 请求
}

func newCgiBinComponentApiGetAuthorizerInfoResult(result CgiBinComponentApiGetAuthorizerInfoResponse, body []byte, http gorequest.Response) *CgiBinComponentApiGetAuthorizerInfoResult {
	return &CgiBinComponentApiGetAuthorizerInfoResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentApiGetAuthorizerInfo 获取授权帐号详情
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/api_get_authorizer_info.html
func (c *Client) CgiBinComponentApiGetAuthorizerInfo(ctx context.Context) (*CgiBinComponentApiGetAuthorizerInfoResult, error) {
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
	param := gorequest.NewParams()
	param["component_appid"] = c.GetComponentAppId()   // 第三方平台 appid
	param["authorizer_appid"] = c.GetAuthorizerAppid() // 授权方 appid
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/component/api_get_authorizer_info?component_access_token=%v", c.GetComponentAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response CgiBinComponentApiGetAuthorizerInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newCgiBinComponentApiGetAuthorizerInfoResult(response, request.ResponseBody, request), nil
}
