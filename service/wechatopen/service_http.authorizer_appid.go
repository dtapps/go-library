package wechatopen

import (
	"context"
	"fmt"
	"net/http"
)

// ServeHttpAuthorizerAppid 授权跳转
func (c *Client) ServeHttpAuthorizerAppid(ctx context.Context, w http.ResponseWriter, r *http.Request) (resp CgiBinComponentApiQueryAuthResponse, agentUserId string, err error) {

	var (
		query = r.URL.Query()

		authCode  = query.Get("auth_code")
		expiresIn = query.Get("expires_in")
	)

	agentUserId = query.Get("agent_user_id")

	if authCode == "" {
		return resp, agentUserId, fmt.Errorf("找不到授权码参数")
	}

	if expiresIn == "" {
		return resp, agentUserId, fmt.Errorf("找不到过期时间参数")
	}

	info, err := c.CgiBinComponentApiQueryAuth(ctx, authCode)
	if err != nil {
		return resp, agentUserId, err
	}

	if info.AuthorizationInfo.AuthorizerAppid == "" {
		return resp, agentUserId, fmt.Errorf("获取失败")
	}

	return info, agentUserId, nil
}
