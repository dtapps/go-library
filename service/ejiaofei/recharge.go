package ejiaofei

import (
	"context"
	"encoding/xml"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"net/http"
)

type RechargeResponse struct {
	XMLName xml.Name `xml:"response"`
	UserID  string   `xml:"userid"`  // 用户账号
	OrderID string   `xml:"orderid"` // 用户提交订单号
	Face    float64  `xml:"face"`    // 官方价格
	Price   float64  `xml:"price"`   // 用户成本价
	Error   int64    `xml:"error"`   // 错误提示
}

type RechargeResult struct {
	Result RechargeResponse   // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newRechargeResult(result RechargeResponse, body []byte, http gorequest.Response) *RechargeResult {
	return &RechargeResult{Result: result, Body: body, Http: http}
}

// Recharge 充值接口
// rechargeType = 充值类型（1-话费，2-流量，3-加油卡） 由鼎信商务提供
// orderId = 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
// account = 充值账户 需要充值的手机号码
// face = 充值面值 话费以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信; 加油卡以元为单位，包含50、100、200、500、800、1000、2000; 流量待定
// isSlowRecharge = 是否慢充 0-否 1-是 默认不传会指定快充
// area = 充值流量范围 0 全国流量，1 省内流量
// effectTime = 生效日期 0 即时生效，1次日生效，2 次月生效
// validity = 流量有效期数量 所传为正整数
// unit = 流量有效期 0/小时，1/天 ，2/个月，3/季度，4/年
// oilType = 加油卡类型（300-中石化，310-中石油）
// operator = 运营商
func (c *Client) Recharge(ctx context.Context, rechargeType int64, orderId string, account string, face int64, notMustParams ...gorequest.Params) (*RechargeResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("rechargeType", rechargeType)              // 充值类型（1-话费，2-流量，3-加油卡） 由鼎信商务提供
	params.Set("orderId", orderId)                        // 用户提交的订单号 用户提交的订单号，最长32位（用户保证其唯一性）
	params.Set("appId", c.GetUserId())                    // 用户编号 由鼎信商务提供
	params.Set("account", account)                        // 充值账户 需要充值的手机号码
	params.Set("timeStamp", gotime.Current().Timestamp()) // 时间戳 时间戳(北京时间)，只允许当前服务器时间正负五分钟内的请求
	params.Set("appSecret", c.GetPwd())                   // 加密密码 由鼎信商务提供
	params.Set("face", face)                              // 充值面值 话费以元为单位，包含10、20、30、50、100、200、300、500 移动联通电信; 加油卡以元为单位，包含50、100、200、500、800、1000、2000; 流量待定
	// 请求
	request, err := c.requestJson(ctx, apiUrl+"/recharge.do", params, http.MethodGet)
	if err != nil {
		return newRechargeResult(RechargeResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RechargeResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newRechargeResult(response, request.ResponseBody, request), err
}
