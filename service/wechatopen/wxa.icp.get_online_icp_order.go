package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetOnlineIcpOrderResponse struct {
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
	IcpSubject struct {
		BaseInfo struct {
			Type         int    `json:"type"`          // 主体性质，示例值：5
			Name         string `json:"name"`          // 主办单位名称，示例值："张三"
			Province     string `json:"province"`      // 备案省份，使用省份代码，示例值："110000"(参考：获取区域信息接口)
			City         string `json:"city"`          // 备案城市，使用城市代码，示例值："110100"(参考：获取区域信息接口)
			District     string `json:"district"`      // 备案县区，使用县区代码，示例值："110105"(参考：获取区域信息接口)
			Address      string `json:"address"`       // 通讯地址，必须属于备案省市区，地址开头的省市区不用填入，例如：通信地址为“北京市朝阳区高碑店路181号1栋12345室”时，只需要填写 "高碑店路181号1栋12345室" 即可
			Comment      string `json:"comment"`       // 主体信息备注，根据需要，如实填写
			RecordNumber string `json:"record_number"` // 主体备案号，示例值：粤B2-20090059（申请小程序备案时不用填写，查询已备案详情时会返回）
		} `json:"base_info"` // 主体基本信息
		PersonalInfo struct {
			ResidencePermit string `json:"residence_permit"` // 临时居住证明照片 media_id，个人备案且非本省人员，需要提供居住证、暂住证、社保证明、房产证等临时居住证明，示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
		} `json:"personal_info"` // 个人主体额外信息
		OrganizeInfo struct {
			CertificateType    int    `json:"certificate_type"`    // 主体证件类型，示例值：2(参考：获取证件类型接口)
			CertificateNumber  string `json:"certificate_number"`  // 主体证件号码，示例值："110105199001011234"
			CertificateAddress string `json:"certificate_address"` // 主体证件住所，示例值："北京市朝阳区高碑店路181号1栋12345室"
			CertificatePhoto   string `json:"certificate_photo"`   // 主体证件照片 media_id，如果小程序主体为非个人类型，则必填，示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
		} `json:"organize_info"` // 主体额外信息（个人备案时，如果存在与主体负责人信息相同的字段，则填入相同的值）
		PrincipalInfo struct {
			Name                         string `json:"name"`                            // 负责人姓名，示例值："张三"
			Mobile                       string `json:"mobile"`                          // 负责人联系方式，示例值："13012344321"
			Email                        string `json:"email"`                           // 负责人电子邮件，示例值："zhangsan@zhangsancorp.com"
			EmergencyContact             string `json:"emergency_contact"`               // 负责人应急联系方式，示例值："17743211234"
			CertificateType              int    `json:"certificate_type"`                // 负责人证件类型，示例值：2(参考：获取证件类型接口，此处只能填入单位性质属于个人的证件类型)
			CertificateNumber            string `json:"certificate_number"`              // 负责人证件号码，示例值："110105199001011234"
			CertificateValidityDateStart string `json:"certificate_validity_date_start"` // 负责人证件有效期起始日期，格式为 YYYYmmdd，示例值："20230815"
			CertificateValidityDateEnd   string `json:"certificate_validity_date_end"`   // 负责人证件有效期终止日期，格式为 YYYYmmdd，如证件长期有效，请填写 "长期"，示例值："20330815"
			CertificatePhotoFront        string `json:"certificate_photo_front"`         // 负责人证件正面照片 media_id（身份证为人像面），示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			CertificatePhotoBack         string `json:"certificate_photo_back"`          // 负责人证件背面照片 media_id（身份证为国徽面），示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			AuthorizationLetter          string `json:"authorization_letter"`            // 授权书 media_id，当主体负责人不是法人时需要主体负责人授权书，当小程序负责人不是法人时需要小程序负责人授权书，示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			VerifyTaskId                 string `json:"verify_task_id"`                  // 扫脸认证任务id(扫脸认证接口返回的task_id)，仅小程序负责人需要扫脸，主体负责人无需扫脸，示例值："R5PqRPNb6GmG3i0rqd4pTg"
		} `json:"principal_info"` // 主体负责人信息
		LegalPersonInfo struct {
			Name              string `json:"name"`               // 法人姓名，示例值："张三"
			CertificateNumber string `json:"certificate_number"` // 法人证件号码，示例值："110105199001011234"
		} `json:"legal_person_info"` // 法人信息（非个人备案，且主体负责人不是法人时，必填）
	} `json:"icp_subject"` // 备案主体信息，不包括图片、视频材料(参考：申请小程序备案接口的 ICPSubject)
	IcpApplets struct {
		BaseInfo struct {
			Appid               string  `json:"appid"`                 // 小程序ID，不用填写，后台自动拉取
			Name                string  `json:"name"`                  // 	小程序名称，不用填写，后台自动拉取
			ServiceContentTypes []int64 `json:"service_content_types"` // 小程序服务内容类型，只能填写二级服务内容类型，最多5个，示例值：[3, 4](参考：获取小程序服务类型接口)
			NrlxDetails         []struct {
				Type int    `json:"type"` // 前置审批类型，示例值：2(参考：获取前置审批项接口)
				Cpde string `json:"cpde"` // 前置审批号，如果前置审批类型不是“以上都不涉及”，则必填，示例值："粤-12345号"
				Name string `json:"name"` // 前置审批媒体材料 media_id，如果前置审批类型不是“以上都不涉及”，则必填，示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			} `json:"nrlx_details"` // 前置审批项，列表中不能存在重复的前置审批类型id，如不涉及前置审批项，也需要填“以上都不涉及”
			Comment      string `json:"comment"`       // 请具体描述小程序实际经营内容、主要服务内容，该信息为主管部门审核重要依据，备注内容字数限制20-200字，请认真填写。（特殊备注要求请查看注意事项）
			RecordNumber string `json:"record_number"` // 小程序备案号，示例值：粤B2-20090059-1626X（申请小程序备案时不用填写，查询已备案详情时会返回）
		} `json:"base_info"` // 	微信小程序基本信息
		PrincipalInfo struct {
			Name                         string `json:"name"`                            // 负责人姓名，示例值："张三"
			Mobile                       string `json:"mobile"`                          // 负责人联系方式，示例值："13012344321"
			Email                        string `json:"email"`                           // 负责人电子邮件，示例值："zhangsan@zhangsancorp.com"
			EmergencyContact             string `json:"emergency_contact"`               // 负责人应急联系方式，示例值："17743211234"
			CertificateType              int    `json:"certificate_type"`                // 负责人证件类型，示例值：2(参考：获取证件类型接口，此处只能填入单位性质属于个人的证件类型)
			CertificateNumber            string `json:"certificate_number"`              // 负责人证件号码，示例值："110105199001011234"
			CertificateValidityDateStart string `json:"certificate_validity_date_start"` // 负责人证件有效期起始日期，格式为 YYYYmmdd，示例值："20230815"
			CertificateValidityDateEnd   string `json:"certificate_validity_date_end"`   // 负责人证件有效期终止日期，格式为 YYYYmmdd，如证件长期有效，请填写 "长期"，示例值："20330815"
			CertificatePhotoFront        string `json:"certificate_photo_front"`         // 负责人证件正面照片 media_id（身份证为人像面），示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			CertificatePhotoBack         string `json:"certificate_photo_back"`          // 负责人证件背面照片 media_id（身份证为国徽面），示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			AuthorizationLetter          string `json:"authorization_letter"`            // 授权书 media_id，当主体负责人不是法人时需要主体负责人授权书，当小程序负责人不是法人时需要小程序负责人授权书，示例值："4ahCGpd3CYkE6RpkNkUR5czt3LvG8xDnDdKAz6bBKttSfM8p4k5Rj6823HXugPwQBurgMezyib7"
			VerifyTaskId                 string `json:"verify_task_id"`                  // 扫脸认证任务id(扫脸认证接口返回的task_id)，仅小程序负责人需要扫脸，主体负责人无需扫脸，示例值："R5PqRPNb6GmG3i0rqd4pTg"
		} `json:"principal_info"` // 小程序负责人信息
	} `json:"icp_applets"` // 微信小程序信息，不包括图片、视频材料(参考：申请小程序备案接口的 ICPApplets)
}

type GetOnlineIcpOrderResult struct {
	Result GetOnlineIcpOrderResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newGetOnlineIcpOrderResult(result GetOnlineIcpOrderResponse, body []byte, http gorequest.Response) *GetOnlineIcpOrderResult {
	return &GetOnlineIcpOrderResult{Result: result, Body: body, Http: http}
}

// GetOnlineIcpOrder 获取小程序已备案详情
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/record/getOnlineIcpOrder.html
func (c *Client) GetOnlineIcpOrder(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetOnlineIcpOrderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetOnlineIcpOrderResponse
	request, err := c.request(ctx, "wxa/icp/get_online_icp_order?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetOnlineIcpOrderResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetOnlineIcpOrderResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 86328:
		return "无法找到资源"
	default:
		return resp.Result.Errmsg
	}
}
