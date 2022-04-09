package request

type TaobaoTopSdkFeedbackUploadRequest struct {
	/*
	   1、回传加密信息     */
	Type *string `json:"type" required:"true" `
	/*
	   具体内容，json形式     */
	Content *string `json:"content,omitempty" required:"false" `
}

func (s *TaobaoTopSdkFeedbackUploadRequest) SetType(v string) *TaobaoTopSdkFeedbackUploadRequest {
	s.Type = &v
	return s
}
func (s *TaobaoTopSdkFeedbackUploadRequest) SetContent(v string) *TaobaoTopSdkFeedbackUploadRequest {
	s.Content = &v
	return s
}

func (req *TaobaoTopSdkFeedbackUploadRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Type != nil {
		paramMap["type"] = *req.Type
	}
	if req.Content != nil {
		paramMap["content"] = *req.Content
	}
	return paramMap
}

func (req *TaobaoTopSdkFeedbackUploadRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
