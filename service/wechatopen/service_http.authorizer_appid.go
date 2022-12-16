package wechatopen

import (
	"context"
	"errors"
	"net/http"
)

// ServeHttpAuthorizerAppid 授权跳转
func (c *Client) ServeHttpAuthorizerAppid(ctx context.Context, r *http.Request) (resp CgiBinComponentApiQueryAuthResponse, agentUserId string, err error) {
	var (
		query = r.URL.Query()

		authCode  = query.Get("auth_code")
		expiresIn = query.Get("expires_in")
	)

	agentUserId = query.Get("agent_user_id")

	if authCode == "" {
		return resp, agentUserId, errors.New("找不到授权码参数")
	}

	if expiresIn == "" {
		return resp, agentUserId, errors.New("找不到过期时间参数")
	}

	info, err := c.CgiBinComponentApiQueryAuth(ctx, authCode)
	if err != nil {
		return resp, agentUserId, err
	}
	if info.Result.AuthorizationInfo.AuthorizerAppid == "" {
		return resp, agentUserId, errors.New("获取失败")
	}

	return info.Result, agentUserId, nil
}
