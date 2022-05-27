package gitee

import (
	"encoding/json"
	"fmt"
	"time"
)

// UserResult 返回参数
type UserResult struct {
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

// User 获取授权用户的资料 https://gitee.com/api/v5/swagger#/getV5User
func (app *App) User() (result UserResult, err error) {
	// request
	body, err := app.request(fmt.Sprintf("https://gitee.com/api/v5/user?access_token=%s", app.AccessToken), map[string]interface{}{}, "GET")

	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
