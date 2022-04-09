package request

type TaobaoTbkItemInfoGetRequest struct {
	/*
	   商品ID串，用,分割，最大40个     */
	NumIids *string `json:"num_iids" required:"true" `
	/*
	   链接形式：1：PC，2：无线，默认：１ defalutValue��1    */
	Platform *int64 `json:"platform,omitempty" required:"false" `
	/*
	   ip地址，影响邮费获取，如果不传或者传入不准确，邮费无法精准提供     */
	Ip *string `json:"ip,omitempty" required:"false" `
}

func (s *TaobaoTbkItemInfoGetRequest) SetNumIids(v string) *TaobaoTbkItemInfoGetRequest {
	s.NumIids = &v
	return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetPlatform(v int64) *TaobaoTbkItemInfoGetRequest {
	s.Platform = &v
	return s
}
func (s *TaobaoTbkItemInfoGetRequest) SetIp(v string) *TaobaoTbkItemInfoGetRequest {
	s.Ip = &v
	return s
}

func (req *TaobaoTbkItemInfoGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.NumIids != nil {
		paramMap["num_iids"] = *req.NumIids
	}
	if req.Platform != nil {
		paramMap["platform"] = *req.Platform
	}
	if req.Ip != nil {
		paramMap["ip"] = *req.Ip
	}
	return paramMap
}

func (req *TaobaoTbkItemInfoGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
