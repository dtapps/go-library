package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinComponentFastRegisterWeAppCreateResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type CgiBinComponentFastRegisterWeAppCreateResult struct {
	Result CgiBinComponentFastRegisterWeAppCreateResponse // 结果
	Body   []byte                                         // 内容
	Http   gorequest.Response                             // 请求
}

func newCgiBinComponentFastRegisterWeAppCreateResult(result CgiBinComponentFastRegisterWeAppCreateResponse, body []byte, http gorequest.Response) *CgiBinComponentFastRegisterWeAppCreateResult {
	return &CgiBinComponentFastRegisterWeAppCreateResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentFastRegisterWeAppCreate 快速注册企业小程序
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/register-management/fast-registration-ent/registerMiniprogram.html
func (c *Client) CgiBinComponentFastRegisterWeAppCreate(ctx context.Context, componentAccessToken string, notMustParams ...gorequest.Params) (*CgiBinComponentFastRegisterWeAppCreateResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/component/fastregisterweapp?action=create&component_access_token="+componentAccessToken, params, http.MethodPost)
	if err != nil {
		return newCgiBinComponentFastRegisterWeAppCreateResult(CgiBinComponentFastRegisterWeAppCreateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinComponentFastRegisterWeAppCreateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinComponentFastRegisterWeAppCreateResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *CgiBinComponentFastRegisterWeAppCreateResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 89249:
		return "该 appid 已有转正任务执行中，距上次任务 24h 后再试"
	case 89247:
		return "系统内部错误"
	case 86004:
		return "无效微信号"
	case 61070:
		return "法人姓名与微信号不一致"
	case 89248:
		return "企业代码类型无效，请选择正确类型填写"
	case 89250:
		return "未找到该任务"
	case 89251:
		return "模板消息已下发，待法人人脸核身校验"
	case 89252:
		return "法人&企业信息一致性校验中"
	case 89253:
		return "缺少参数"
	case 89254:
		return "第三方权限集不全，请补充权限集后重试"
	case 89255:
		return "code参数无效，请检查 code 长度以及内容是否正确；注意code_type的值不同需要传的 code 长度不一样"
	}
	return "系统繁忙"
}

// StatusInfo 状态描述
func (resp *CgiBinComponentFastRegisterWeAppCreateResult) StatusInfo(status int) string {
	switch status {
	case 100001:
		return "已下发的模板消息法人并未确认且已超时（24h），未进行身份证校验"
	case 100002:
		return "已下发的模板消息法人并未确认且已超时（24h），未进行人脸识别校验"
	case 100003:
		return "已下发的模板消息法人并未确认且已超时（24h）"
	case 101:
		return "工商数据返回：“企业已注销”"
	case 102:
		return "工商数据返回：“企业不存在或企业信息未更新”"
	case 103:
		return "工商数据返回：“企业法定代表人姓名不一致”"
	case 104:
		return "工商数据返回：“企业法定代表人身份证号码不一致”"
	case 105:
		return "法定代表人身份证号码，工商数据未更新，请 5-15 个工作日之后尝试"
	case 1000:
		return "工商数据返回：“企业信息或法定代表人信息不一致”"
	case 1001:
		return "主体创建小程序数量达到上限"
	case 1002:
		return "主体违规命中黑名单"
	case 1003:
		return "管理员绑定账号数量达到上限"
	case 1004:
		return "管理员违规命中黑名单"
	case 1005:
		return "管理员手机绑定账号数量达到上限"
	case 1006:
		return "管理员手机号违规命中黑名单"
	case 1007:
		return "管理员身份证创建账号数量达到上限"
	case 1008:
		return "管理员身份证违规命中黑名单"
	case -1:
		return "企业与法人姓名不一致"
	}
	return fmt.Sprintf("%v", status)
}
