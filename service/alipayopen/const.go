package alipayopen

const (
	apiUrl = "https://openapi.alipay.com/gateway.do"
)

const (
	LogTable = "alipayopen"
)

const AuthorizationCode = "authorization_code" // 表示换取使用用户授权码code换取授权令牌
const RefreshToken = "refresh_token"           // 表示使用refresh_token刷新获取新授权令牌

type ApiError struct {
	ErrorResponse struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"`
	Sign string `json:"sign"`
}
