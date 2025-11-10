package wechatopen

type APIResponse struct {
	Errcode int    `json:"errcode,omitempty"` // 错误码
	Errmsg  string `json:"errmsg,omitempty"`  // 错误信息
}

type APIRetResponse struct {
	Ret    int    `json:"ret,omitempty"`     // 错误码
	ErrMsg string `json:"err_msg,omitempty"` // 错误信息
}
