package gitee

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
	"time"
)

type ApiV5UserResponse struct {
	Id                int64     `json:"id"`
	Login             string    `json:"login"`
	Name              string    `json:"name"`
	AvatarUrl         string    `json:"avatar_url"`
	Url               string    `json:"url"`
	HtmlUrl           string    `json:"html_url"`
	FollowersUrl      string    `json:"followers_url"`
	FollowingUrl      string    `json:"following_url"`
	GistsUrl          string    `json:"gists_url"`
	StarredUrl        string    `json:"starred_url"`
	SubscriptionsUrl  string    `json:"subscriptions_url"`
	OrganizationsUrl  string    `json:"organizations_url"`
	ReposUrl          string    `json:"repos_url"`
	EventsUrl         string    `json:"events_url"`
	ReceivedEventsUrl string    `json:"received_events_url"`
	Type              string    `json:"type"`
	Blog              string    `json:"blog"`
	Weibo             string    `json:"weibo"`
	Bio               string    `json:"bio"`
	PublicRepos       int       `json:"public_repos"`
	PublicGists       int       `json:"public_gists"`
	Followers         int       `json:"followers"`
	Following         int       `json:"following"`
	Stared            int       `json:"stared"`
	Watched           int       `json:"watched"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Email             string    `json:"email"`
}

type ApiV5UserResult struct {
	Result ApiV5UserResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newApiV5UserResult(result ApiV5UserResponse, body []byte, http gorequest.Response, err error) *ApiV5UserResult {
	return &ApiV5UserResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiV5User 获取授权用户的资料
// https://gitee.com/api/v5/swagger#/getV5User
func (c *Client) ApiV5User(accessToken string) *ApiV5UserResult {
	// 参数
	params := gorequest.NewParamsWith()
	// 请求
	request, err := c.request(apiUrl+fmt.Sprintf("/api/v5/user?access_token=%s", accessToken), params, http.MethodGet)
	// 定义
	var response ApiV5UserResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newApiV5UserResult(response, request.ResponseBody, request, err)
}
