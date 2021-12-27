package taobao

type TbkActivityInfoGetResult struct {
	TbkActivityInfoGetResponse struct {
		Data struct {
			WxQrcodeUrl       string `json:"wx_qrcode_url"`
			ClickUrl          string `json:"click_url"`
			ShortClickUrl     string `json:"short_click_url"`
			TerminalType      string `json:"terminal_type"`
			MaterialOssUrl    string `json:"material_oss_url"`
			PageName          string `json:"page_name"`
			PageStartTime     string `json:"page_start_time"`
			PageEndTime       string `json:"page_end_time"`
			WxMiniprogramPath string `json:"wx_miniprogram_path"`
		} `json:"data"`
	} `json:"tbk_activity_info_get_response"`
}

// TbkActivityInfoGet 淘宝客-推广者-官方活动转链 https://open.taobao.com/api.htm?spm=a219a.7386797.0.0.5a83669a7rURsF&source=search&docId=48340&docType=2
func (app *App) TbkActivityInfoGet(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := NewParamsWithType("taobao.tbk.activity.info.get", notMustParams...)
	// 请求
	body, err = app.request(params)
	return
}