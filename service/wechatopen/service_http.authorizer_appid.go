package wechatopen

import (
	"errors"
	"net/http"
	"strconv"
)

// ServeHttpAuthorizerAppid 授权跳转
func (app *App) ServeHttpAuthorizerAppid(r *http.Request) (resp CgiBinComponentApiQueryAuthResponse, agentUserId int64, pacId uint, err error) {
	var (
		query = r.URL.Query()

		authCode  = query.Get("auth_code")
		expiresIn = query.Get("expires_in")
	)

	agentUserId = ToInt64(query.Get("agent_user_id"))

	pacId = ToUint(query.Get("pac_id"))

	if authCode == "" {
		return resp, agentUserId, pacId, errors.New("找不到授权码参数")
	}

	if expiresIn == "" {
		return resp, agentUserId, pacId, errors.New("找不到过期时间参数")
	}

	info := app.CgiBinComponentApiQueryAuth(authCode)
	if info.Result.AuthorizationInfo.AuthorizerAppid == "" {
		return resp, agentUserId, pacId, errors.New("获取失败")
	}

	return info.Result, agentUserId, pacId, nil
}

// ToFloat64 string到float64
func ToFloat64(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

// ToInt64 string到int64
func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	return int64(ToFloat64(s))
}

// ToUint string到uint64
func ToUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return uint(i)
	}
	return 0
}
