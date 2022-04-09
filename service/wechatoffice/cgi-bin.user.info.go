package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CgiBinUserInfoResponse struct {
	Subscribe      int           `json:"subscribe"`       // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息
	Openid         string        `json:"openid"`          // 用户的标识，对当前公众号唯一
	Language       string        `json:"language"`        // 用户的语言，简体中文为zh_CN
	SubscribeTime  int           `json:"subscribe_time"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	Unionid        string        `json:"unionid"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段
	Remark         string        `json:"remark"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	Groupid        int           `json:"groupid"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      []interface{} `json:"tagid_list"`      // 用户被打上的标签ID列表
	SubscribeScene string        `json:"subscribe_scene"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_REPRINT 他人转载 ,ADD_SCENE_LIVESTREAM 视频号直播，ADD_SCENE_CHANNELS 视频号 , ADD_SCENE_OTHERS 其他
	QrScene        int           `json:"qr_scene"`        // 二维码扫码场景（开发者自定义）
	QrSceneStr     string        `json:"qr_scene_str"`    // 二维码扫码场景描述（开发者自定义）
	Nickname       string        `json:"nickname"`        // 用户昵称
	Sex            int           `json:"sex"`             // 	用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province       string        `json:"province"`        // 用户个人资料填写的省份
	City           string        `json:"city"`            // 普通用户个人资料填写的城市
	Country        string        `json:"country"`         // 国家，如中国为CN
	Headimgurl     string        `json:"headimgurl"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效
}
type CgiBinUserInfoResult struct {
	Result CgiBinUserInfoResponse // 结果
	Body   []byte                 // 内容
	Err    error                  // 错误
}

func NewCgiBinUserInfoResult(result CgiBinUserInfoResponse, body []byte, err error) *CgiBinUserInfoResult {
	return &CgiBinUserInfoResult{Result: result, Body: body, Err: err}
}

// CgiBinUserInfo 获取用户基本信息(UnionID机制)
// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
func (app *App) CgiBinUserInfo(openid string) *CgiBinUserInfoResult {
	app.AccessToken = app.GetAccessToken()
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN", app.AccessToken, openid), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinUserInfoResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinUserInfoResult(response, body, err)
}
