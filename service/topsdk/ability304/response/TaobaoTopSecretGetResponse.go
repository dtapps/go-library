package response

type TaobaoTopSecretGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   下次更新秘钥间隔，单位（秒）
	*/
	Interval int64 `json:"interval,omitempty" `
	/*
	   最长有效期，容灾使用，单位（秒）
	*/
	MaxInterval int64 `json:"max_interval,omitempty" `
	/*
	   秘钥值
	*/
	Secret string `json:"secret,omitempty" `
	/*
	   秘钥版本号
	*/
	SecretVersion int64 `json:"secret_version,omitempty" `
	/*
	   app配置信息
	*/
	AppConfig string `json:"app_config,omitempty" `
}
