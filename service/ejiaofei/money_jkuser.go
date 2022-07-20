package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type MoneyJkUserResponse struct {
	XMLName   xml.Name `xml:"response"`
	LastMoney float64  `xml:"lastMoney"` // 用户余额
	Tag       int      `xml:"tag"`       // 用户状态（0正常 1暂停）
	Error     int      `xml:"error"`     // 错误提示
}

type MoneyJkUserResult struct {
	Result MoneyJkUserResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func newMoneyJkUserResult(result MoneyJkUserResponse, body []byte, http gorequest.Response, err error) *MoneyJkUserResult {
	return &MoneyJkUserResult{Result: result, Body: body, Http: http, Err: err}
}

// MoneyJkUser 用户余额查询
func (c *Client) MoneyJkUser() *MoneyJkUserResult {
	// 签名
	c.signStr = fmt.Sprintf("userid%vpwd%v", c.getUserId(), c.getPwd())
	// 请求
	request, err := c.request(apiUrl+"/money_jkuser.do", map[string]interface{}{}, http.MethodGet)
	// 定义
	var response MoneyJkUserResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newMoneyJkUserResult(response, request.ResponseBody, request, err)
}
