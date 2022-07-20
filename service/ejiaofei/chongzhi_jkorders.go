package ejiaofei

import (
	"encoding/xml"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ChOngZhiJkOrdersResponse struct {
	XMLName   xml.Name `xml:"response"`
	UserID    string   `xml:"userid"`    // 会员账号
	PorderID  string   `xml:"Porderid"`  // 鼎信平台订单号
	OrderID   string   `xml:"orderid"`   // 用户订单号
	Account   string   `xml:"account"`   // 需要充值的手机号码
	Face      string   `xml:"face"`      // 充值面值
	Amount    string   `xml:"amount"`    // 购买数量
	StartTime string   `xml:"starttime"` // 开始时间
	State     string   `xml:"state"`     // 订单状态
	EndTime   string   `xml:"endtime"`   // 结束时间
	Error     string   `xml:"error"`     // 错误提示
}

type ChOngZhiJkOrdersResult struct {
	Result ChOngZhiJkOrdersResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
	Err    error                    // 错误
}

func newChOngZhiJkOrdersResult(result ChOngZhiJkOrdersResponse, body []byte, http gorequest.Response, err error) *ChOngZhiJkOrdersResult {
	return &ChOngZhiJkOrdersResult{Result: result, Body: body, Http: http, Err: err}
}

// ChOngZhiJkOrders 话费充值接口
// orderid 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
// face 充值面值	以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信
// account 手机号码	需要充值的手机号码
func (c *Client) ChOngZhiJkOrders(orderID string, face int, account string) *ChOngZhiJkOrdersResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("orderid", orderID)
	param.Set("face", face)
	param.Set("account", account)
	param.Set("amount", 1)
	params := gorequest.NewParamsWith(param)
	// 签名
	c.signStr = fmt.Sprintf("userid%vpwd%vorderid%vface%vaccount%vamount1", c.getUserId(), c.getPwd(), orderID, face, account)
	// 请求
	request, err := c.request(apiUrl+"/chongzhi_jkorders.do", params, http.MethodGet)
	// 定义
	var response ChOngZhiJkOrdersResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newChOngZhiJkOrdersResult(response, request.ResponseBody, request, err)
}
