package wechatoffice

// Result 接口
type Result struct {
	Byte                  []byte // 内容
	Err                   error  // 错误
	GetCallBackIpResponse        // IP即微信调用开发者服务器所使用的出口IP
}
