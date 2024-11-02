package ejiaofei

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gorequest"
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
// orderid = 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
// face = 充值面值 以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信
// account = 手机号码 需要充值的手机号码
// amount = 购买数量 只能为1
// operator = 运营商可指定当前手机号的运营商信息进行充值,为空则自动匹配号段对应的运营商进行充值; 具体对应的运营商信息表3.3
func (c *Client) ChOngZhiJkOrders(ctx context.Context, orderid string, face int64, account string, amount int64, notMustParams ...gorequest.Params) (*ChOngZhiJkOrdersResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("userid", c.GetUserId()) // 用户编号
	params.Set("pwd", c.GetPwd())       // 加密密码
	params.Set("orderid", orderid)      // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	params.Set("face", face)            // 充值面值 以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信
	params.Set("account", account)      // 手机号码 需要充值的手机号码
	params.Set("amount", amount)        // 购买数量 只能为1

	// 响应
	var response ChOngZhiJkOrdersResponse

	// 请求
	request, err := c.requestXml(ctx, "chongzhi_jkorders.do", params, http.MethodGet, &response)
	return newChOngZhiJkOrdersResult(response, request.ResponseBody, request), err
}
