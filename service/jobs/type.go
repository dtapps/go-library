package jobs

const (
	// 微信支付查询
	TypeWechatPayPrepaidRefill = "wechat.pay.prepaid_refill" //【话费充值】
	TypeWechatPayPower         = "wechat.pay.power"          //【电费充值】
	TypeWechatPayKashangwl     = "wechat.pay.kashangwl"      //【特惠充值】
	TypeWechatPayEastiot       = "wechat.pay.eastiot"        //【东方物联】
	TypeWechatPayEjiaofei      = "wechat.pay.ejiaofei"       //【易缴费】
	TypeWechatPayMovie         = "wechat.pay.movie"          //【守兔】

	// 接口支付申请
	TypeApiPaySubmitPrepaidRefill = "api.pay.submit.prepaid_refill" //【话费充值】
	TypeApiPaySubmitPower         = "api.pay.submit.power"          //【电费充值】
	TypeApiPaySubmitKashangwl     = "api.pay.submit.kashangwl"      //【特惠充值】
	TypeApiPaySubmitEastiot       = "api.pay.submit.eastiot"        //【东方物联】
	TypeApiPaySubmitEjiaofei      = "api.pay.submit.ejiaofei"       //【易缴费】
	TypeApiPaySubmitMovie         = "api.pay.submit.movie"          //【守兔】

	// 接口支付查询
	TypeApiPayQueryPrepaidRefill = "api.pay.query.prepaid_refill" //【话费充值】
	TypeApiPayQueryPower         = "api.pay.query.power"          //【电费充值】
	TypeApiPayQueryKashangwl     = "api.pay.query.kashangwl"      //【特惠充值】
	TypeApiPayQueryEastiot       = "api.pay.query.eastiot"        //【东方物联】
	TypeApiPayQueryEjiaofei      = "api.pay.query.ejiaofei"       //【易缴费】
	TypeApiPayQueryMovie         = "api.pay.query.movie"          //【守兔】
	TypeApiPayQueryPinduoduo     = "api.pay.query.pinduoduo"      //【拼多多】
	TypeApiPayQueryMeituan       = "api.pay.query.meituan"        //【美团】

	// 微信支付退款申请
	TypeWechatRefundsSubmitPrepaidRefill = "wechat.refunds.submit.prepaid_refill" //【话费充值】
	TypeWechatRefundsSubmitPower         = "wechat.refunds.submit.power"          //【电费充值】
	TypeWechatRefundsSubmitKashangwl     = "wechat.refunds.submit.kashangwl"      //【特惠充值】
	TypeWechatRefundsSubmitEastiot       = "wechat.refunds.submit.eastiot"        //【东方物联】
	TypeWechatRefundsSubmitEjiaofei      = "wechat.refunds.submit.ejiaofei"       //【易缴费】
	TypeWechatRefundsSubmitMovie         = "wechat.refunds.submit.movie"          //【守兔】

	// 微信支付退款查询
	TypeWechatRefundsQueryPrepaidRefill = "wechat.refunds.query.prepaid_refill" //【话费充值】
	TypeWechatRefundsQueryPower         = "wechat.refunds.query.power"          //【电费充值】
	TypeWechatRefundsQueryKashangwl     = "wechat.refunds.query.kashangwl"      //【特惠充值】
	TypeWechatRefundsQueryEastiot       = "wechat.refunds.query.eastiot"        //【东方物联】
	TypeWechatRefundsQueryEjiaofei      = "wechat.refunds.query.ejiaofei"       //【易缴费】
	TypeWechatRefundsQueryMovie         = "wechat.refunds.query.movie"          //【守兔】

	// 返拥
	TypeGoldenBeansIssuePrepaidRefill = "golden_beans.issue.prepaid_refill" //【话费充值】
	TypeGoldenBeansIssuePower         = "golden_beans.issue.power"          //【电费充值】
	TypeGoldenBeansIssueKashangwl     = "golden_beans.issue.kashangwl"      //【特惠充值】
	TypeGoldenBeansIssueEastiot       = "golden_beans.issue.eastiot"        //【东方物联】
	TypeGoldenBeansIssueMovie         = "golden_beans.issue.movie"          //【守兔】
	TypeGoldenBeansIssueRewardedvideo = "golden_beans.issue.rewardedvideo"  //【激励视频】

	// TypeGoldenBeansFrozenPinduoduo 【拼多多】冻结返拥金豆
	TypeGoldenBeansFrozenPinduoduo = "golden_beans.frozen.pinduoduo"
	// TypeGoldenBeansFrozenMeituan 【美团】冻结返拥金豆
	TypeGoldenBeansFrozenMeituan = "golden_beans.frozen.meituan"

	// TypeGoldenBeansThawSuccessPinduoduo 【拼多多】解冻返拥金豆正常发放
	TypeGoldenBeansThawSuccessPinduoduo = "golden_beans.thaw.success.pinduoduo"
	// TypeGoldenBeansThawSuccessMeituan 【美团】解冻返拥金豆正常发放
	TypeGoldenBeansThawSuccessMeituan = "golden_beans.thaw.success.meituan"

	// TypeGoldenBeansThawErrorPinduoduo 【拼多多】解冻返拥金豆扣款
	TypeGoldenBeansThawErrorPinduoduo = "golden_beans.thaw.error.pinduoduo"
	// TypeGoldenBeansThawErrorMeituan 【美团】解冻返拥金豆扣款
	TypeGoldenBeansThawErrorMeituan = "golden_beans.thaw.error.meituan"

	// 抵扣金豆退款
	TypeGoldenBeansRefundsPrepaidRefill = "golden_beans.refunds.prepaid_refill" //【话费充值】
	TypeGoldenBeansRefundsPower         = "golden_beans.refunds.power"          //【电费充值】

	// 金豆退款
	TypeGoldenBeansRefundsEjiaofei = "golden_beans.refunds.ejiaofei" //【易缴费】

	// 客服进度
	TypeCustomerAutoPrepaidRefill = "customer.auto.prepaid_refill" //【话费充值】
	TypeCustomerAutoPower         = "customer.auto.power"          //【电费充值】
	TypeCustomerAutoKashangwl     = "customer.auto.kashangwl"      //【特惠充值】
	TypeCustomerAutoEjiaofei      = "customer.auto.ejiaofei"       //【易缴费】
	TypeCustomerAutoMovie         = "customer.auto.movie"          //【守兔】

	// 订单同步
	TypeSyncOrderPinduoduo = "sync.order.pinduoduo" //【拼多多】
	TypeSyncOrderMeituan   = "sync.order.meituan"   //【美团】

	// 商品同步
	TypeSyncGoodsEastiot              = "sync.goods.eastiot"                //【东方物联】
	TypeSyncGoodsTypeEastiot          = "sync.goods.type.eastiot"           //【东方物联】
	TypeSyncGoodsPriceAllKashangwl    = "sync.goods.price.all.kashangwl"    //【特惠充值】
	TypeSyncGoodsPriceSingleKashangwl = "sync.goods.price.single.kashangwl" //【特惠充值】
	TypeCheckGoodsAllPinduoduo        = "check.goods.all.pinduoduo"         //【拼多多】
	TypeCheckGoodsSinglePinduoduo     = "check.goods.single.pinduoduo"      //【拼多多】

	// 微信更新
	TypeSyncWechat = "sync.wechat" //【微信】token/ticket同步

	// 商家金豆转换
	TypeGoldenBeanConversionUserGoldenBean       = "admin.golden_bean.conversion.user.golden_bean"        //【商家金豆】转用户金豆
	TypeGoldenBeanConversionUserDeductGoldenBean = "admin.golden_bean.conversion.user.deduct.golden_bean" //【商家金豆】转用户抵扣金豆

	// 企业自定义
	TypeNewServiceAgentIssue     = "new_service.agent.issue"      //【企业自定义】【代理商】下发检查
	TypeNewServiceAgentIssueNext = "new_service.agent.issue.next" //【企业自定义】【代理商】下发到商家

	// 团队
	TypeTeamInv = "team.inv" //【团队邀请】

	// 修复商家账号数量
	TypeRepairMerchantAccountQuantityTesting = "repair_merchant.account.quantity.testing" //【修复】【商家账号】【数量】查询出等级进行下一步
	TypeRepairMerchantAccountQuantityLevel   = "repair_merchant.account.quantity.level"   //【修复】【商家账号】【数量】通过等级来修复

	// 观察接口
	TypeApiPayObservationPrepaidRefill    = "api.pay.observation.prepaid_refill"    //{600}【话费充值】
	TypeApiPayObservationPower            = "api.pay.observation.power"             //{600}【电费充值】
	TypeApiPayObservationKashangwl        = "api.pay.observation.kashangwl"         //{600}【特惠充值】
	TypeApiPayObservationEjiaofei         = "api.pay.observation.ejiaofei"          //{600}【易缴费】
	TypeApiPayObservationMovie            = "api.pay.observation.movie"             //{600}【守兔】
	TypeApiPayObservationPinduoduo        = "api.pay.observation.pinduoduo"         //{600}【拼多多】
	TypeApiPayObservationMeituan          = "api.pay.observation.meituan"           //{600}【美团】
	TypeApiPayObservationWechat           = "api.pay.observation.wechat"            //{600}【微信支付】
	TypeApiPayObservationWechatRefunds    = "api.pay.observation.wechat.refunds"    //{600}【微信支付退款】
	TypeApiPayObservationWechatWithdrawal = "api.pay.observation.wechat.withdrawal" //{600}【微信支付提现】

	// 微信提现
	TypeWechatPayWithdrawalSubmit  = "wechat.pay.withdrawal.submit"  // 微信支付提现申请
	TypeWechatPayWithdrawalQuery   = "wechat.pay.withdrawal.query"   // 微信支付提现查询
	TypeWechatPayWithdrawalRefunds = "wechat.pay.withdrawal.refunds" // 微信支付提现退款
)

func GetTypeWechatPay(Type string) string {
	return "wechat.pay." + Type
}

func GetTypeApiPaySubmit(Type string) string {
	return "api.pay.submit." + Type
}

func GetTypeApiPayQuery(Type string) string {
	return "api.pay.query." + Type
}

func GetTypeWechatRefundsSubmit(Type string) string {
	return "wechat.refunds.submit." + Type
}

func GetTypeWechatRefundsQuery(Type string) string {
	return "wechat.refunds.query." + Type
}

func GetTypeGoldenBeansIssue(Type string) string {
	return "golden_beans.issue." + Type
}

func GetTypeGoldenBeansRefunds(Type string) string {
	return "golden_beans.refunds." + Type
}

func GetTypeCustomerAuto(Type string) string {
	return "customer.auto." + Type
}
