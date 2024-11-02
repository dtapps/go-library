package pinduoduo

import (
	"context"
	"errors"
	"net/http"
)

// ServeHttpAuthorizer 授权跳转
func (c *Client) ServeHttpAuthorizer(ctx context.Context, w http.ResponseWriter, r *http.Request) (string, string, error) {

	var (
		query = r.URL.Query()

		code  = query.Get("code")
		state = query.Get("state")
	)

	if code == "" {
		err := errors.New("找不到授权码参数")
		return code, state, err
	}

	return code, state, nil
}
