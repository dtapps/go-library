package request

type TaobaoTopSecretGetRequest struct {
	/*
	   秘钥版本号     */
	SecretVersion *int64 `json:"secret_version,omitempty" required:"false" `
	/*
	   伪随机数     */
	RandomNum *string `json:"random_num" required:"true" `
	/*
	   自定义用户id     */
	CustomerUserId *int64 `json:"customer_user_id,omitempty" required:"false" `
}

func (s *TaobaoTopSecretGetRequest) SetSecretVersion(v int64) *TaobaoTopSecretGetRequest {
	s.SecretVersion = &v
	return s
}
func (s *TaobaoTopSecretGetRequest) SetRandomNum(v string) *TaobaoTopSecretGetRequest {
	s.RandomNum = &v
	return s
}
func (s *TaobaoTopSecretGetRequest) SetCustomerUserId(v int64) *TaobaoTopSecretGetRequest {
	s.CustomerUserId = &v
	return s
}

func (req *TaobaoTopSecretGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.SecretVersion != nil {
		paramMap["secret_version"] = *req.SecretVersion
	}
	if req.RandomNum != nil {
		paramMap["random_num"] = *req.RandomNum
	}
	if req.CustomerUserId != nil {
		paramMap["customer_user_id"] = *req.CustomerUserId
	}
	return paramMap
}

func (req *TaobaoTopSecretGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
