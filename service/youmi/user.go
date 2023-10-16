package youmi

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type UserResponse struct {
	Errno  int64  `json:"errno"`  // 错误码，0代表成功，非0代表失败
	Errmsg string `json:"errmsg"` // 错误描述
	Data   struct {
		Id       int64  `json:"id"`       // userid
		Username string `json:"username"` // 名称
		Balance  string `json:"balance"`  // 余额
	} `json:"data,omitempty"`
}

type UserResult struct {
	Result UserResponse       // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newUserResult(result UserResponse, body []byte, http gorequest.Response) *UserResult {
	return &UserResult{Result: result, Body: body, Http: http}
}

// User 查询用户信息
// https://www.showdoc.com.cn/dyr/9227004018562421
func (c *Client) User(ctx context.Context, notMustParams ...gorequest.Params) (*UserResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserID()) // 账号ID
	// 请求
	request, err := c.request(ctx, c.GetApiURL()+"index/user", params)
	if err != nil {
		return newUserResult(UserResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response UserResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newUserResult(response, request.ResponseBody, request), err
}
