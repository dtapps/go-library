package v2

import "encoding/json"

type GetVersionResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

// GetVersion 获取同步版本号 https://www.showdoc.com.cn/1154868044931571/6566701084841699
func (app *App) GetVersion() (result GetVersionResult, err error) {
	body, err := app.request("movieapi/movie-info/get-version", map[string]interface{}{})
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
