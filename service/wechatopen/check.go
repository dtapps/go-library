package wechatopen

import (
	"context"
	"errors"
)

func (c *Client) checkAuthorizerIsConfig(ctx context.Context) error {
	authorizerAppid := c.GetAuthorizerAppid(ctx)
	if authorizerAppid == "" {
		return errors.New("请配置 authorizerAppid")
	}
	return nil
}

func (c *Client) checkAuthorizerConfig(ctx context.Context) error {
	authorizerAppid := c.GetAuthorizerAppid(ctx)
	if authorizerAppid == "" {
		return errors.New("请配置 authorizerAppid")
	}
	authorizerAccessToken := c.GetAuthorizerAccessToken(ctx)
	if authorizerAccessToken == "" {
		return errors.New("请配置 authorizerAccessToken")
	}
	return nil
}
