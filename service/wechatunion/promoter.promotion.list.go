package wechatunion

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PromotionListResponse struct {
	Errcode             int    `json:"errcode"` // 错误码
	Errmsg              string `json:"errmsg"`  // 错误信息
	PromotionSourceList []struct {
		PromotionSourceName string `json:"promotionSourceName"` // 推广位名称
		PromotionSourcePid  string `json:"promotionSourcePid"`  // 推广位ID，PID
		Status              string `json:"status"`              // 状态
		PidId               string `json:"pidId"`
	} `json:"promotionSourceList"` // 推广位数据
	Total           int `json:"total"`           // 推广位总数
	PromotionMaxCnt int `json:"promotionMaxCnt"` // 允许创建的推广位最大数量
}

type PromotionListResult struct {
	Result PromotionListResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
	Err    error                 // 错误
}

func NewPromotionListResult(result PromotionListResponse, body []byte, http gorequest.Response, err error) *PromotionListResult {
	return &PromotionListResult{Result: result, Body: body, Http: http, Err: err}
}

// PromotionList 获取推广位列表
// https://developers.weixin.qq.com/doc/ministore/union/access-guidelines/promoter/api/promotion.html#_4-%E8%8E%B7%E5%8F%96%E6%8E%A8%E5%B9%BF%E4%BD%8D%E5%88%97%E8%A1%A8
func (c *Client) PromotionList(start int, limit int) *PromotionListResult {
	// 参数
	params := NewParams()
	params.Set("start", start) // 偏移
	params.Set("limit", limit) // 每页条数
	// 请求
	request, err := c.request(UnionUrl+fmt.Sprintf("/promoter/promotion/list?access_token%s", c.getAccessToken()), params, http.MethodGet)
	// 定义
	var response PromotionListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewPromotionListResult(response, request.ResponseBody, request, err)
}
