package pinduoduo

import (
	"context"
	"errors"
	"net/http"
)

// ServeHttpAuthorizerApi 授权跳转
func (c *Client) ServeHttpAuthorizerApi(ctx context.Context, w http.ResponseWriter, r *http.Request) (PopAuthTokenCreate, string, string, error) {

	var (
		query = r.URL.Query()

		code  = query.Get("code")
		state = query.Get("state")
	)

	if code == "" {
		err := errors.New("找不到授权码参数")
		return PopAuthTokenCreate{}, code, state, err
	}

	response, err := c.PopAuthTokenCreate(ctx, code)
	if err != nil {
		return response, code, state, err
	}

	return response, code, state, nil
}
