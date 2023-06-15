package wechatopen

// MonitorAuthorizerAccessToken 授权方access_token 监控
//func MonitorAuthorizerAccessToken(ctx context.Context, c *Client, authorizerRefreshToken string) (string, error) {
//	authorizerAccessToken := GetAuthorizerAccessToken(ctx, c) // 查询
//	// 判断
//	result, err := c.CgiBinGetCallBackIp(ctx, authorizerAccessToken)
//	if err != nil {
//		return "", err
//	}
//	if len(result.Result.IpList) > 0 {
//		return authorizerAccessToken, err
//	}
//	// 重新获取
//	resp, err := c.CgiBinComponentApiAuthorizerToken(ctx, authorizerRefreshToken)
//	if resp.Result.AuthorizerRefreshToken == "" {
//		return authorizerAccessToken, err
//	}
//	return SetAuthorizerAccessToken(ctx, c, resp.Result.AuthorizerAccessToken), err
//}
