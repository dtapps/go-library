package wechatunion

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type PromotionAddResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Pid     string `json:"pid"` // 推广位ID，PID
}

// PromotionAdd 添加推广位 https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html
func (app *App) PromotionAdd(promotionSourceName string) (result PromotionAddResult, err error) {

	if len(app.AccessToken) <= 0 {
		return result, errors.New("调用凭证异常")
	}

	if len(promotionSourceName) > 20 {
		return result, errors.New(fmt.Sprintf("推广位名称（最长20个字，名称不可重复）：%d", len(promotionSourceName)))
	}

	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/promotion/add?access_token%s", app.AccessToken), map[string]interface{}{
		"promotionSourceName": promotionSourceName,
	}, http.MethodPost)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	if result.Errcode == 14005 {
		return result, errors.New(fmt.Sprintf("推广位数量达到上限：%s", err))
	}
	if result.Errcode == 14007 {
		return result, errors.New(fmt.Sprintf("推广位名称重复：%s", err))
	}
	return result, err
}

type PromotionDel struct {
	PromotionSourcePid  string `json:"promotionSourcePid"`  // 推广位PID
	PromotionSourceName string `json:"promotionSourceName"` // 推广位名称
}

type PromotionDelResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// PromotionDel 删除某个推广位 https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html
func (app *App) PromotionDel(param PromotionDel) (result PromotionDelResult, err error) {

	if len(app.AccessToken) <= 0 {
		return result, errors.New("调用凭证异常")
	}

	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}

	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/promotion/del?access_token%s", app.AccessToken), params, http.MethodPost)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, err

}

type PromotionUpd struct {
	PreviousPromotionInfo struct {
		PromotionSourcePid  string `json:"promotionSourcePid"`  // 要修改的推广位PID
		PromotionSourceName string `json:"promotionSourceName"` // 修改前名称
	} `json:"previousPromotionInfo"`
	PromotionInfo struct {
		PromotionSourceName string `json:"promotionSourceName"` // 修改后名称
	} `json:"promotionInfo"`
}

type PromotionUpdResult struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// PromotionUpd 修改指定的推广位名称 https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html
func (app *App) PromotionUpd(param PromotionUpd) (result PromotionUpdResult, err error) {

	if len(app.AccessToken) <= 0 {
		return result, errors.New("调用凭证异常")
	}

	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}

	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/promotion/upd?access_token%s", app.AccessToken), params, http.MethodPost)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, err

}

type PromotionListResult struct {
	Errcode             int    `json:"errcode"`
	Errmsg              string `json:"errmsg"`
	PromotionSourceList []struct {
		PromotionSourceName string `json:"promotionSourceName"` // 推广位名称
		PromotionSourcePid  string `json:"promotionSourcePid"`  // 推广位ID，PID
		Status              string `json:"status"`              // 状态
		PidId               string `json:"pidId"`
	} `json:"promotionSourceList"` // 推广位数据
	Total           int `json:"total"`           // 推广位总数
	PromotionMaxCnt int `json:"promotionMaxCnt"` // 允许创建的推广位最大数量
}

// PromotionList 获取推广位列表 https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html
func (app *App) PromotionList(start int, limit int) (result PromotionListResult, err error) {

	if len(app.AccessToken) <= 0 {
		return result, errors.New("调用凭证异常")
	}

	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/union/promoter/promotion/list?access_token%s", app.AccessToken), map[string]interface{}{
		"start": start, // 偏移
		"limit": limit, // 每页条数
	}, http.MethodGet)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, err

}
