package jobs

func GetTypeApiPaySubmit(Type string) string {
	return "api.pay.submit." + Type
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
