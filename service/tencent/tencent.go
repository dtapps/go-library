package tencent

type ConfigTencent struct {
	SecretId  string
	SecretKey string
}

type Tencent struct {
	config     ConfigTencent
	Lighthouse lighthouse
}

func NewTencent(config *ConfigTencent) *Tencent {
	t := &Tencent{}
	t.config.SecretId = config.SecretId
	t.config.SecretKey = config.SecretKey
	return t
}
