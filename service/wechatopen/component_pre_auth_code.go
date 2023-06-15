package wechatopen

// MonitorComponentPreAuthCode 第三方平台预授权码 监控
//func MonitorComponentPreAuthCode(ctx context.Context, c *Client) (string, error) {
//	// 查询
//	preAuthCode := GetComponentPreAuthCode(ctx, c)
//	// 判断
//	if preAuthCode != "" {
//		return preAuthCode, nil
//	}
//	// 重新获取
//	resp, err := c.CgiBinComponentApiCreatePreAuthCoden(ctx)
//	if resp.Result.PreAuthCode == "" {
//		return preAuthCode, err
//	}
//	return SetComponentPreAuthCode(ctx, c, resp.Result.PreAuthCode), err
//}
