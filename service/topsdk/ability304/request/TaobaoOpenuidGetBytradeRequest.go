package request

type TaobaoOpenuidGetBytradeRequest struct {
	/*
	   订单ID     */
	Tid *int64 `json:"tid" required:"true" `
}

func (s *TaobaoOpenuidGetBytradeRequest) SetTid(v int64) *TaobaoOpenuidGetBytradeRequest {
	s.Tid = &v
	return s
}

func (req *TaobaoOpenuidGetBytradeRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Tid != nil {
		paramMap["tid"] = *req.Tid
	}
	return paramMap
}

func (req *TaobaoOpenuidGetBytradeRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
