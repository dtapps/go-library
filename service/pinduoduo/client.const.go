package pinduoduo

const (
	apiUrl = "https://gw-api.pinduoduo.com/api/router"
)

const (
	LogTable = "pinduoduo"
)

type ApiErrorT struct {
	ErrorResponse struct {
		ErrorMsg  string `json:"error_msg"`
		SubMsg    string `json:"sub_msg"`
		SubCode   string `json:"sub_code"`
		ErrorCode int    `json:"error_code"`
		RequestId string `json:"request_id"`
	} `json:"error_response"`
}
