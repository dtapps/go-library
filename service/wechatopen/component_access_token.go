package wechatopen

// MonitorComponentAccessToken 第三方平台access_token 监控
//func MonitorComponentAccessToken(ctx context.Context, c *Client) (string, error) {
//	componentAccessToken := GetComponentAccessToken(ctx, c) // 查询
//	// 判断
//	result, err := c.CgiBinGetApiDomainIp(ctx, componentAccessToken)
//	if err != nil {
//		return "", err
//	}
//	if len(result.Result.IpList) > 0 {
//		return componentAccessToken, err
//	}
//	// 重新获取
//	resp, err := c.CgiBinComponentApiComponentToken(ctx)
//	if resp.Result.ComponentAccessToken == "" {
//		return componentAccessToken, err
//	}
//	return SetComponentAccessToken(ctx, c, resp.Result.ComponentAccessToken), err
//}
