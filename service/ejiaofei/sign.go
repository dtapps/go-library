package ejiaofei

import (
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"sort"
)

func (c *Client) xmlSign(url string, param *gorequest.Params) (signStr string) {
	switch url {

	case "http://api.ejiaofei.net:11140/checkCost.do":
		// 会员订单成本价查询接口
		signStr = fmt.Sprintf("userid%vpwd%vorderid%v", c.GetUserId(), c.GetPwd(), param.Get("orderid"))

	case "http://api.ejiaofei.net:11140/chongzhi_jkorders.do":
		// 话费充值接口
		signStr = fmt.Sprintf("userid%vpwd%vorderid%vface%vaccount%vamount%v", c.GetUserId(), c.GetPwd(), param.Get("orderid"), param.Get("face"), param.Get("account"), param.Get("amount"))

	case "http://api.ejiaofei.net:11140/gprsChongzhiAdvance.do":
		// 流量充值接口
		signStr = fmt.Sprintf("userid%vpwd%vorderid%vaccount%vgprs%varea%veffecttime%vvalidity%vtimes%v", c.GetUserId(), c.GetPwd(), param.Get("orderid"), param.Get("account"), param.Get("gprs"), param.Get("area"), param.Get("effecttime"), param.Get("validity"), param.Get("times"))

	case "http://api.ejiaofei.net:11140/money_jkuser.do":
		// 用户余额查询
		signStr = fmt.Sprintf("userid%vpwd%v", c.GetUserId(), c.GetPwd())

	case "http://api.ejiaofei.net:11140/query_jkorders.do":
		// 通用查询接口
		signStr = fmt.Sprintf("userid%vpwd%vorderid%v", c.GetUserId(), c.GetPwd(), param.Get("orderid"))

	case "http://api.ejiaofei.net:11140/queryTXproduct.do":
		// 可充值腾讯产品查询
		signStr = fmt.Sprintf("userid%vpwd%v", c.GetUserId(), c.GetPwd())

	case "http://api.ejiaofei.net:11140/txchongzhi.do":
		// 流量充值接口
		signStr = fmt.Sprintf("userid%vpwd%vorderid%vaccount%vproductid%vamount%vip%vtimes%v", c.GetUserId(), c.GetPwd(), param.Get("orderid"), param.Get("account"), param.Get("productid"), param.Get("amount"), param.Get("ip"), param.Get("times"))

	}

	return gomd5.ToUpper(fmt.Sprintf("%s%s", signStr, c.GetKey()))
}

func (c *Client) jsonSign(param *gorequest.Params) string {
	var keys []string
	for k := range param.DeepGet() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s%s", key, gostring.ToString(param.Get(key)))
	}
	signStr += fmt.Sprintf("%s", c.GetKey())
	return gomd5.ToUpper(signStr)
}
