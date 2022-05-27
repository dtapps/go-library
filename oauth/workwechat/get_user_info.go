package workwechat

import (
	"encoding/json"
	"fmt"
)

type GetUserInfo struct {
	Code string `json:"code"`
}

// GetUserInfoResult 返回参数
type GetUserInfoResult struct {
	Errcode        int    `json:"errcode"`         // 返回码
	Errmsg         string `json:"errmsg"`          // 对返回码的文本描述内容
	UserId         string `json:"userId"`          // 成员UserID。若需要获得用户详情信息
	OpenId         string `json:"OpenId"`          // 非企业成员的标识，对当前企业唯一。不超过64字节
	DeviceId       string `json:"DeviceId"`        // 手机设备号(由企业微信在安装时随机生成，删除重装会改变，升级不受影响)
	ExternalUserid string `json:"external_userid"` // 外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回。如果是第三方应用调用，针对同一个客户，同一个服务商不同应用获取到的id相同
}

// GetUserInfo 获取访问用户身份 https://open.work.weixin.qq.com/api/doc/90000/90135/91023
func (app *App) GetUserInfo(param GetUserInfo) (result GetUserInfoResult, err error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s", app.AccessToken, param.Code)

	// request
	body, err := app.request(url, map[string]interface{}{}, "GET")
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
