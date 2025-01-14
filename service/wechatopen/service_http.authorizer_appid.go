package wechatopen

import (
	"context"
	"errors"
	"net/http"
)

// ServeHttpAuthorizerAppid 授权跳转
func (c *Client) ServeHttpAuthorizerAppid(ctx context.Context, w http.ResponseWriter, r *http.Request, componentAccessToken string) (resp CgiBinComponentApiQueryAuthResponse, agentUserId string, err error) {

	var (
		query = r.URL.Query()

		authCode  = query.Get("auth_code")
		expiresIn = query.Get("expires_in")
	)

	agentUserId = query.Get("agent_user_id")

	if authCode == "" {
		err = errors.New("找不到授权码参数")
		return resp, agentUserId, err
	}

	if expiresIn == "" {
		err = errors.New("找不到过期时间参数")
		return resp, agentUserId, err
	}

	info, err := c.CgiBinComponentApiQueryAuth(ctx, componentAccessToken, authCode)
	if err != nil {
		return resp, agentUserId, err
	}

	if info.Result.AuthorizationInfo.AuthorizerAppid == "" {
		err = errors.New("获取失败")
		return resp, agentUserId, err
	}

	return info.Result, agentUserId, nil
}
