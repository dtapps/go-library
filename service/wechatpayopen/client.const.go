package wechatpayopen

const (
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n" // 数字签名原文格式
	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "WECHATPAY2-SHA256-RSA2048 mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
)

func getAuthorizationType() string {
	return "WECHATPAY2-" + algorithm()
}

// 返回使用的签名算法：SHA256-RSA2048
func algorithm() string {
	return "SHA256-RSA2048"
}

// 接口状态
const (
	CodeSUCCESS    = "SUCCESS"    // 支付成功 退款成功
	CodeREFUND     = "REFUND"     // 转入退款
	CodeNOTPAY     = "NOTPAY"     // 未支付
	CodeCLOSED     = "CLOSED"     // 已关闭 退款关闭
	CodeREVOKED    = "REVOKED"    // 已撤销
	CodeUSERPAYING = "USERPAYING" // 用户支付中
	CodePAYERROR   = "PAYERROR"   // 支付失败
	CodePROCESSING = "PROCESSING" // 退款处理中
	CodeABNORMAL   = "ABNORMAL"   // 退款异常
)

type ApiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

const (
	MERCHANT_ID         = "MERCHANT_ID"
	PERSONAL_OPENID     = "PERSONAL_OPENID"
	PERSONAL_SUB_OPENID = "PERSONAL_SUB_OPENID"
)
