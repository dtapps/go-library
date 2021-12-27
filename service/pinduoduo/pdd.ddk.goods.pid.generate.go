package pinduoduo

// GoodsPidGenerateResult 返回参数
type GoodsPidGenerateResult struct {
	PIdGenerateResponse struct {
		PIdList []struct {
			CreateTime int    `json:"create_time,omitempty"` // 推广位创建时间
			PidName    string `json:"pid_name,omitempty"`    // 推广位名称
			PId        string `json:"p_id,omitempty"`        // 调用方推广位ID
			MediaId    int    `json:"media_id,omitempty"`    // 媒体id
		} `json:"p_id_list"`
		RemainPidCount int `json:"remain_pid_count"` // PID剩余数量
	} `json:"p_id_generate_response"`
}

// GoodsPidGenerate 创建多多进宝推广位 https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.goods.pid.generate
func (app *App) GoodsPidGenerate(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := NewParamsWithType("pdd.ddk.goods.pid.generate", notMustParams...)
	// 请求
	body, err = app.request(params)
	return
}
