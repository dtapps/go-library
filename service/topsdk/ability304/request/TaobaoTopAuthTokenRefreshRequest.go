package request

type TaobaoTopAuthTokenRefreshRequest struct {
	/*
	   grantType==refresh_token 时需要     */
	RefreshToken *string `json:"refresh_token" required:"true" `
}

func (s *TaobaoTopAuthTokenRefreshRequest) SetRefreshToken(v string) *TaobaoTopAuthTokenRefreshRequest {
	s.RefreshToken = &v
	return s
}

func (req *TaobaoTopAuthTokenRefreshRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefreshToken != nil {
		paramMap["refresh_token"] = *req.RefreshToken
	}
	return paramMap
}

func (req *TaobaoTopAuthTokenRefreshRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
