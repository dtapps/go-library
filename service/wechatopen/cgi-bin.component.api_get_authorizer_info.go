package wechatopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiGetAuthorizerInfoResponse struct {
	AuthorizerInfo struct {
		NickName        string `json:"nick_name"` // 昵称
		HeadImg         string `json:"head_img"`  // 头像
		ServiceTypeInfo struct {
			Id   int64  `json:"id"`   // 类型id
			Name string `json:"name"` // 类型说明
		} `json:"service_type_info"` // 小程序类型
		VerifyTypeInfo struct {
			Id   int64  `json:"id"`   // 类型id
			Name string `json:"name"` // 类型说明
		} `json:"verify_type_info"` // 小程序认证类型
		UserName     string `json:"user_name"`       // 原始 ID
		Alias        string `json:"alias,omitempty"` // 公众号所设置的微信号，可能为空
		QrcodeUrl    string `json:"qrcode_url"`      // 二维码图片的 URL，开发者最好自行也进行保存
		BusinessInfo struct {
			OpenPay   int64 `json:"open_pay"`   // 是否开通微信支付功能
			OpenShake int64 `json:"open_shake"` // 是否开通微信摇一摇功能
			OpenScan  int64 `json:"open_scan"`  // 是否开通微信扫商品功能
			OpenCard  int64 `json:"open_card"`  // 是否开通微信卡券功能
			OpenStore int64 `json:"open_store"` // 是否开通微信门店功能
		} `json:"business_info"` // 用以了解功能的开通状况（0代表未开通，1代表已开通)
		Idc             int64  `json:"idc,omitempty"`
		PrincipalName   string `json:"principal_name"` // 主体名称
		Signature       string `json:"signature"`      // 帐号介绍
		MiniProgramInfo struct {
			Network struct {
				RequestDomain   []string `json:"RequestDomain"`   // request合法域名
				WsRequestDomain []string `json:"WsRequestDomain"` // socket合法域名
				UploadDomain    []string `json:"UploadDomain"`    // uploadFile合法域名
				DownloadDomain  []string `json:"DownloadDomain"`  // downloadFile合法域名
				UDPDomain       []string `json:"UDPDomain"`       // udp合法域名
				TCPDomain       []any    `json:"TCPDomain"`       // tcp合法域名
			} `json:"network"` // 小程序配置的合法域名信息
			Categories []struct {
				First  string `json:"first"`  // 一级类目
				Second string `json:"second"` // 二级类目
			} `json:"categories"` // 小程序配置的类目信息
			VisitStatus int64 `json:"visit_status,omitempty"`
		} `json:"MiniProgramInfo"` // 小程序配置，根据这个字段判断是否为小程序类型授权
		RegisterType  int64 `json:"register_type"`  // 小程序注册方式
		AccountStatus int64 `json:"account_status"` // 帐号状态，该字段小程序也返回
		BasicConfig   struct {
			IsPhoneConfigured bool `json:"is_phone_configured"` // 是否已经绑定手机号
			IsEmailConfigured bool `json:"is_email_configured"` // 是否已经绑定邮箱，不绑定邮箱帐号的不可登录微信公众平台
		} `json:"basic_config"` // 基础配置信息
	} `json:"authorizer_info"` // 小程序帐号信息
	AuthorizationInfo struct {
		AuthorizerAppid        string `json:"authorizer_appid"`         // 授权的公众号或者小程序 appid
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"` // 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效
		FuncInfo               []struct {
			FuncscopeCategory struct {
				Id   int64  `json:"id"`   // 权限集id
				Type int64  `json:"type"` // 权限集类型
				Name string `json:"name"` // 权限集名称
				Desc string `json:"desc"` // 权限集描述
			} `json:"funcscope_category"` // 	授权给开发者的权限集详情
		} `json:"func_info"` // 授权给第三方平台的权限集id列表，权限集id的含义可查看权限集介绍
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
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/authorization-management/getAuthorizerInfo.html
func (c *Client) CgiBinComponentApiGetAuthorizerInfo(ctx context.Context, authorizerAppid, componentAccessToken string, notMustParams ...*gorequest.Params) (*CgiBinComponentApiGetAuthorizerInfoResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId()) // 第三方平台appid
	params.Set("authorizer_appid", authorizerAppid)      // 授权方appid

	// 请求
	var response CgiBinComponentApiGetAuthorizerInfoResponse
	request, err := c.request(ctx, fmt.Sprintf("cgi-bin/component/api_get_authorizer_info?access_token=%s", componentAccessToken), params, http.MethodPost, &response)
	return newCgiBinComponentApiGetAuthorizerInfoResult(response, request.ResponseBody, request), err
}
