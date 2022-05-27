package jobs

import "go.dtapp.net/library/utils/goip"

var TypeIp = map[string]string{
	// 微信支付查询
	TypeWechatPayPrepaidRefill: "",
	TypeWechatPayPower:         "",
	TypeWechatPayKashangwl:     "",
	TypeWechatPayEastiot:       "",
	TypeWechatPayEjiaofei:      "",
	TypeWechatPayMovie:         "",

	// 接口支付申请
	TypeApiPaySubmitPrepaidRefill: "",
	TypeApiPaySubmitPower:         "",
	TypeApiPaySubmitKashangwl:     "",
	TypeApiPaySubmitEastiot:       "47.112.146.193",
	TypeApiPaySubmitEjiaofei:      "47.112.146.193",
	TypeApiPaySubmitMovie:         "",

	// 接口支付查询
	TypeApiPayQueryPrepaidRefill: "",
	TypeApiPayQueryPower:         "",
	TypeApiPayQueryKashangwl:     "",
	TypeApiPayQueryEastiot:       "47.112.146.193",
	TypeApiPayQueryEjiaofei:      "",
	TypeApiPayQueryMovie:         "",
	TypeApiPayQueryPinduoduo:     "",
	TypeApiPayQueryMeituan:       "",

	// 微信支付退款申请
	TypeWechatRefundsSubmitPrepaidRefill: "",
	TypeWechatRefundsSubmitPower:         "",
	TypeWechatRefundsSubmitKashangwl:     "",
	TypeWechatRefundsSubmitEastiot:       "",
	TypeWechatRefundsSubmitEjiaofei:      "",
	TypeWechatRefundsSubmitMovie:         "",

	// 微信支付退款查询
	TypeWechatRefundsQueryPrepaidRefill: "",
	TypeWechatRefundsQueryPower:         "",
	TypeWechatRefundsQueryKashangwl:     "",
	TypeWechatRefundsQueryEastiot:       "",
	TypeWechatRefundsQueryEjiaofei:      "",
	TypeWechatRefundsQueryMovie:         "",

	// 返拥
	TypeGoldenBeansIssuePrepaidRefill: "",
	TypeGoldenBeansIssuePower:         "",
	TypeGoldenBeansIssueKashangwl:     "",
	TypeGoldenBeansIssueEastiot:       "",
	TypeGoldenBeansIssueMovie:         "",
	TypeGoldenBeansIssueRewardedvideo: "",

	// 冻结返拥金豆
	TypeGoldenBeansFrozenPinduoduo: "",
	TypeGoldenBeansFrozenMeituan:   "",

	// 解冻返拥金豆正常发放
	TypeGoldenBeansThawSuccessPinduoduo: "",
	TypeGoldenBeansThawSuccessMeituan:   "",

	// 解冻返拥金豆扣款
	TypeGoldenBeansThawErrorPinduoduo: "",
	TypeGoldenBeansThawErrorMeituan:   "",

	// 抵扣金豆退款
	TypeGoldenBeansRefundsPrepaidRefill: "",
	TypeGoldenBeansRefundsPower:         "",

	// 金豆退款
	TypeGoldenBeansRefundsEjiaofei: "",

	// 客服进度
	TypeCustomerAutoPrepaidRefill: "",
	TypeCustomerAutoPower:         "",
	TypeCustomerAutoKashangwl:     "",
	TypeCustomerAutoEjiaofei:      "",
	TypeCustomerAutoMovie:         "",

	// 订单同步
	TypeSyncOrderPinduoduo: "",
	TypeSyncOrderMeituan:   "",

	// 商品同步
	TypeSyncGoodsEastiot:              "47.112.146.193",
	TypeSyncGoodsTypeEastiot:          "47.112.146.193",
	TypeSyncGoodsPriceAllKashangwl:    "",
	TypeSyncGoodsPriceSingleKashangwl: "",
	TypeCheckGoodsAllPinduoduo:        "",
	TypeCheckGoodsSinglePinduoduo:     "",

	// 微信更新
	TypeSyncWechat: "119.3.235.113,119.3.132.197,47.112.146.193",

	// 商家金豆转换
	TypeGoldenBeanConversionUserGoldenBean:       "",
	TypeGoldenBeanConversionUserDeductGoldenBean: "",

	// 企业自定义
	TypeNewServiceAgentIssue:     "",
	TypeNewServiceAgentIssueNext: "",

	// 团队
	TypeTeamInv: "",

	// 修复商家账号数量
	TypeRepairMerchantAccountQuantityTesting: "",
	TypeRepairMerchantAccountQuantityLevel:   "",

	// 观察接口
	TypeApiPayObservationPrepaidRefill:    "",
	TypeApiPayObservationPower:            "",
	TypeApiPayObservationKashangwl:        "",
	TypeApiPayObservationEjiaofei:         "47.112.146.193",
	TypeApiPayObservationMovie:            "",
	TypeApiPayObservationPinduoduo:        "",
	TypeApiPayObservationMeituan:          "",
	TypeApiPayObservationWechat:           "",
	TypeApiPayObservationWechatRefunds:    "",
	TypeApiPayObservationWechatWithdrawal: "",

	// 微信提现
	TypeWechatPayWithdrawalSubmit:  "47.112.146.193",
	TypeWechatPayWithdrawalQuery:   "",
	TypeWechatPayWithdrawalRefunds: "",
}

func (app *App) Ip(Type string) string {
	return TypeIp[Type]
}

func (app *App) RefreshIp() {
	xip := goip.GetOutsideIp()
	if app.OutsideIp == "" || app.OutsideIp == "0.0.0.0" {
		return
	}
	if app.OutsideIp == xip {
		return
	}
	app.Db.Where("ips = ?", app.OutsideIp).Delete(&TaskIp{}) // 删除
	app.OutsideIp = xip
}
