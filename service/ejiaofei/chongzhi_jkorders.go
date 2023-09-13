package ejiaofei

import (
	"context"
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
}

func newChOngZhiJkOrdersResult(result ChOngZhiJkOrdersResponse, body []byte, http gorequest.Response) *ChOngZhiJkOrdersResult {
	return &ChOngZhiJkOrdersResult{Result: result, Body: body, Http: http}
}

// ChOngZhiJkOrders 话费充值接口
// orderID 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
// face 充值面值	以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信
// account 手机号码	需要充值的手机号码
func (c *Client) ChOngZhiJkOrders(ctx context.Context, notMustParams ...gorequest.Params) (*ChOngZhiJkOrdersResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 签名
	c.config.signStr = fmt.Sprintf("userid%vpwd%vorderid%vface%vaccount%vamount1", c.GetUserId(), c.GetPwd(), params["orderid"], params["face"], params["account"], params["amount"])
	// 请求
	request, err := c.request(ctx, apiUrl+"/chongzhi_jkorders.do", params, http.MethodGet)
	if err != nil {
		return newChOngZhiJkOrdersResult(ChOngZhiJkOrdersResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ChOngZhiJkOrdersResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newChOngZhiJkOrdersResult(response, request.ResponseBody, request), err
}
