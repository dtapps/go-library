package wechatopen

// SignDecryptParams 入参
type SignDecryptParams struct {
	Signature    string // 签名串，对应 URL 参数的msg_signature
	Timestamp    string // 时间戳，对应 URL 参数的timestamp
	Nonce        string // 随机串，对应 URL 参数的nonce
	EncryptType  string // 加密类型
	MsgSignature string
	AppId        string // 第三方平台 appid
	Encrypt      string // 加密内容
}
