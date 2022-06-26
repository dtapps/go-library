package gitee

import (
	"fmt"
)

// OauthAuthorize 获取登录地址
func (c *Client) OauthAuthorize() string {
	return fmt.Sprintf(apiUrl+"/oauth/authorize?client_id=%s&redirect_uri=%s&response_type=code", c.config.ClientID, c.config.RedirectUri)
}
