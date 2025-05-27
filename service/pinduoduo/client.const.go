package pinduoduo

const (
	apiUrl = "https://gw-api.pinduoduo.com/api/router"
)

type ApiError struct {
	ErrorResponse struct {
		ErrorMsg  string `json:"error_msg"`
		SubMsg    string `json:"sub_msg"`
		SubCode   string `json:"sub_code"`
		ErrorCode int    `json:"error_code,omitempty"`
		RequestId string `json:"request_id,omitempty"`
	} `json:"error_response"`
}
