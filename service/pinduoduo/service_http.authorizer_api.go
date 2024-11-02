package pinduoduo

import (
	"context"
	"errors"
	"net/http"
)

// ServeHttpAuthorizerApi 授权跳转
func (c *Client) ServeHttpAuthorizerApi(ctx context.Context, w http.ResponseWriter, r *http.Request) (PddDdkPopAuthTokenCreateResponse, string, string, error) {

	var (
		query = r.URL.Query()

		code  = query.Get("code")
		state = query.Get("state")
	)

	if code == "" {
		err := errors.New("找不到授权码参数")
		return PddDdkPopAuthTokenCreateResponse{}, code, state, err
	}

	response, err := c.PopAuthTokenCreate(ctx, code)
	if err != nil {
		return response.Result, code, state, err
	}

	return response.Result, code, state, nil
}
