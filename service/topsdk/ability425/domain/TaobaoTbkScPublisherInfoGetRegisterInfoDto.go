package domain

type TaobaoTbkScPublisherInfoGetRegisterInfoDto struct {
	/*
	   渠道独有 -店铺名称     */
	ShopName *string `json:"shop_name,omitempty" `

	/*
	   渠道独有 -店铺类型     */
	ShopType *string `json:"shop_type,omitempty" `

	/*
	   渠道独有 -手机号码     */
	PhoneNumber *string `json:"phone_number,omitempty" `

	/*
	   渠道独有 -详细地址     */
	DetailAddress *string `json:"detail_address,omitempty" `

	/*
	   渠道独有 -地区     */
	Location *string `json:"location,omitempty" `

	/*
	   渠道独有 -证件类型     */
	ShopCertifyType *string `json:"shop_certify_type,omitempty" `

	/*
	   渠道独有 -对应的证件证件类型编号     */
	CertifyNumber *string `json:"certify_number,omitempty" `

	/*
	   渠道独有 -经营类型     */
	Career *string `json:"career,omitempty" `
}

func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetShopName(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.ShopName = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetShopType(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.ShopType = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetPhoneNumber(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.PhoneNumber = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetDetailAddress(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.DetailAddress = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetLocation(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.Location = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetShopCertifyType(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.ShopCertifyType = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetCertifyNumber(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.CertifyNumber = &v
	return s
}
func (s *TaobaoTbkScPublisherInfoGetRegisterInfoDto) SetCareer(v string) *TaobaoTbkScPublisherInfoGetRegisterInfoDto {
	s.Career = &v
	return s
}
