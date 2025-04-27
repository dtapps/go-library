package meituan

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type MediaQueryOrderResponse struct {
	Status  int    `json:"status"`  // 响应码，0成功，其他失败
	Message string `json:"message"` // 响应文案
	Data    struct {
		ActId    int64  `json:"actId,omitempty"` // 活动物料id，我要推广-活动推广中第一列的id信息
		SkuCount int64  `json:"skuCount"`        // 查询返回本页的数量合计（无实际使用场景，若查询订单购买商品数可以看返回的dataList中skuCount）
		ScrollId string `json:"scrollId"`        // 分页id，当searchType选择2逐页查询时，出参会返回本字段。用于下一页查询的scrollId字段入参使用
		DataList []struct {
			Platform          int64  `json:"platform"`               // 商品所属业务一级分类类型：1 到家及其他业务类型，2 到店业务类型（包含到店美食、休闲生活）
			BusinessLine      int64  `json:"businessLine,omitempty"` // 1）当platform选择到家及其他业务类型时，业务线枚举 1：外卖订单 WAIMAI 2：闪购红包 3：酒旅 4：美团电商订单（团好货） 5：医药 6：拼好饭 7：商品超值券包 COUPON 8：买菜 MAICAI 11：闪购商品 不传则默认传空表示返回除类型7以外的全部类型查询。若输入参数含7 商品超值券包，则只返回商品超值券包订单 2）当platform选择到店业务类型 时，业务线枚举1:到餐 2: 到综
			OrderId           string `json:"orderId"`                // 订单ID
			PayTime           int64  `json:"payTime"`                // 订单支付时间
			PayPrice          string `json:"payPrice"`               // 订单支付价格。针对到餐、到综、酒店、闪购、医药业务类型，为父订单的支付价格，单位元
			UpdateTime        int64  `json:"updateTime"`             // 订单最近一次的更新时间。到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店类型，订单时间为用户买券包的更新时间，非每张券的更新时间。针对以上业务类型，建议查询单张券的更新时间
			CommissionRate    string `json:"commissionRate"`         // 订单预估佣金比例，300表示3%
			Profit            string `json:"profit"`                 // cps类型的预估佣金收入，单位元，1.60表示1.6元
			CpaProfit         string `json:"cpaProfit"`              // cpa类型的预估佣金收入，单位元，6.50表示6.5元
			Sid               string `json:"sid"`                    // 二级媒体身份标识，用于渠道效果追踪
			ProductId         string `json:"productId,omitempty"`    // 产品ID，对应商品查询接口的skuViewId，目前只支持到家外卖商品券、到家医药、到家闪购商品业务、到店业务类型
			ProductName       string `json:"productName"`            // 产品名称，外卖订单展示店铺名称，到店取单个商品券的名称、其他展示全部商品名称
			SpecificationName string `json:"specificationName"`      // 规格信息，同一个商品名称下可以包括不同的规格，对应不同的价格和佣金
			OrderDetail       []struct {
				CouponStatus string `json:"couponStatus,omitempty"` // 本期只有到到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店业务类型展示订单明细，表示商品券/子订单推广计佣状态，1、付款，2、完成（或券已核销），3、结算，4、失效（含取消或风控的情况）
				ItemOrderId  string `json:"itemOrderId,omitempty"`  // 针对到店到餐、到综、酒店商品券，返回商品券的子订单号。其他业务类型不返回
				FinishTime   string `json:"finishTime,omitempty"`   // 1、针对到家外卖商品券，返回商品券核销完成履约的实物菜品订单号对应的完成时间；2、针对到家医药&闪购商品，返回商品订单完成时间；3、针对到店到餐、到综、酒店子订单，返回子订单对应的券核销时间
				BasicAmount  string `json:"basicAmount,omitempty"`  // 商品的计佣金额，每个商品对应的支付分摊金额，单位元
				CouponFee    string `json:"couponFee,omitempty"`    // 商品的佣金，当推广状态为失效、取消、风控时，佣金值为0，单位元
				OrderViewId  string `json:"orderViewId,omitempty"`  // 只对到家外卖商品券有效。商品券的核销完成履约的实物菜品订单号
				RefundAmount string `json:"refundAmount,omitempty"` // 到店到餐、到综、酒店子订单、到家闪购商品、到家医药业务类型的退款金额，到家其他业务类型不返回数据，单位元
				RefundFee    string `json:"refundFee,omitempty"`    // 到店到餐、到综、酒店子订单、到家闪购商品、到家医药业务类型的退款佣金，到家其他业务类型不返回数据，单位元
				RefundTime   string `json:"refundTime,omitempty"`   // 到店到餐、到综、酒店子订单、到家闪购商品、到家医药业务类型的退款时间，到家其他业务类型不返回数据
				SettleTime   string `json:"settleTime,omitempty"`   // 到家商品券/到家闪购商品/到店到餐/到综/酒店子订单的结算时间，完成并且进入结算账期时则变为结算状态。若存在多次结算记录则取最新结算时间
				UpdateTime   string `json:"updateTime,omitempty"`   // 到家商品券/到家闪购商品/到家医药/到店到餐、到综、酒店子订单的更新时间
			} `json:"orderDetail,omitempty"` // 订单详情，只支持到家外卖商品券、到家医药、到家闪购商品业务、到店到餐、到综、酒店类型返回数据
			RefundPrice     string `json:"refundPrice"`            // 只对非到店到餐、非到综、非酒店业务类型有效。订单维度退款价格，该笔订单用户发生退款行为时的退款计佣金额之和，超值券包订单本期不返回退款数据，单位元
			RefundTime      string `json:"refundTime"`             // 只对非到店到餐、非到综、非酒店业务类型有效。订单维度最新一次发生退款的时间；超值券包订单本期不返回退款数据，单位元
			RefundProfit    string `json:"refundProfit"`           // 只对非到店到餐、非到综、非酒店业务类型有效。订单维度退款预估佣金，该笔订单用户发生退款行为时的退款预估佣金金额之和；超值券包订单本期不返回退款数据，单位元
			CpaRefundProfit string `json:"cpaRefundProfit"`        // cpa退款预估佣金，单位元
			Status          string `json:"status"`                 // 表示订单维度状态，枚举有 2：付款（如果是CPA订单则表示奖励已创建） 3：完成 4：取消 5：风控 6：结算。 针对到家商品券订单、到家闪购订单、到家医药订单、到店到餐、到综、酒店业务类型订单则为父订单相关状态，枚举有2：付款，父订单仅付款，至少有任意一个子订单未核销； 3：完成，父订单中所有子订单都核销完成； 4：取消，父订单中的子订单全部退款或过期未使用； 5：风控，父订单中的子订单全部变成风控状态； 6：结算，父订单中所有子订单都结算完成。（CPA订单只有到家闪购订单、到家医药订单、到店到餐、到综业务类型有本状态） 说明： 1、若到店到餐、到综、酒店业务类型父订单、到家闪购商品父订单、到家医药父订单，含有多个状态混合的子订单，则随机取子订单状态作为父订单状态，建议以orderDetail中每张券状态为准 2、含多个商品或券包的订单不建议使用该字段，实际计佣状态以orderDetail中每张券的计佣状态为准。
			TradeType       int64  `json:"tradeType"`              // 交易类型，1：cps，2：cpa
			ActId           int64  `json:"actId,omitempty"`        // 活动物料id，我要推广-活动推广中第一列的id信息
			Appkey          string `json:"appkey"`                 // 归因到的appKey，对应取链时入参的appkey
			SkuCount        int64  `json:"skuCount,omitempty"`     // 表示sku数量，团好货和券包类型的CPS订单返回有值，其余类型订单不返回该值
			CityName        string `json:"cityName,omitempty"`     // 订单所属的城市，目前支持二级城市粒度。目前只支持到家业务类型-商品超值券包业务线。
			CategoryId      int64  `json:"categoryId,omitempty"`   // 订单品类id。
			CategoryName    string `json:"categoryName,omitempty"` // 订单品类名称。
		} `json:"dataList"` // 数据列表
	} `json:"data,omitempty"` // 响应结果信息
}

type MediaQueryOrderResult struct {
	Result MediaQueryOrderResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newMediaQueryOrderResult(result MediaQueryOrderResponse, body []byte, http gorequest.Response) *MediaQueryOrderResult {
	return &MediaQueryOrderResult{Result: result, Body: body, Http: http}
}

// MediaQueryOrder 查询订单接口
// 查询推广的订单明细及佣金信息，包括到店、到家、买菜等业务类型的订单。支持按付款时间或更新时间查询，查询近3个月的订单明细。支持POST方法查询接口。只接受JSON格式。
// https://media.meituan.com/pc/index.html#/materials/api-detail/query_order
// https://page.meituan.net/html/1706169509872_eb0353/index.html
func (c *MediaClient) MediaQueryOrder(ctx context.Context, notMustParams ...*gorequest.Params) (*MediaQueryOrderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response MediaQueryOrderResponse
	request, err := c.request(ctx, "cps_open/common/api/v1/query_order", http.MethodPost, params, &response)
	return newMediaQueryOrderResult(response, request.ResponseBody, request), err
}
