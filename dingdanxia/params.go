package dingdanxia

// Params 请求参数
type Params map[string]interface{}

func NewParams() Params {
	p := make(Params)
	return p
}

func (app *App) NewParamsWith(params ...Params) Params {
	p := make(Params)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}
