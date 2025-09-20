package gosms

import (
	"context"
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 阿里云
type Aliyun struct {
	config openapi.Config           // 配置
	client *dysmsapi20170525.Client // 实例
}

// 初始化
func NewAliyun(ctx context.Context, opts ...Option) (client *Aliyun, err error) {
	options := NewOptions(opts)

	// 初始化
	client = &Aliyun{}

	client.config.AccessKeyId = tea.String(options.accessKeyId)
	client.config.AccessKeySecret = tea.String(options.accessKeySecret)
	client.config.Endpoint = tea.String("dysmsapi.aliyuncs.com")

	client.client, err = dysmsapi20170525.NewClient(&client.config)

	return client, err
}

// Query 查询短信
// bizId = 发送回执ID
// phoneNumber = 接收短信的手机号码
// sendDate = 短信发送日期
// templateParam = 短信模板变量对应的实际值
func (c *Aliyun) Query(ctx context.Context, bizId, phoneNumber string, sendDate string) (err error) {

	// 参数
	querySendDetailsRequest := &dysmsapi20170525.QuerySendDetailsRequest{
		PhoneNumber: tea.String(phoneNumber),
		BizId:       tea.String(bizId),
		SendDate:    tea.String(sendDate),
		CurrentPage: tea.Int64(1),
		PageSize:    tea.Int64(50),
	}
	response, err := c.client.QuerySendDetails(querySendDetailsRequest)
	if err != nil {
		return err
	}

	for _, dto := range response.Body.SmsSendDetailDTOs.SmsSendDetailDTO {
		if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("3"))) {
			return SuccessStatus // 发送成功
		} else if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("2"))) {
			return FailureStatus // 发送失败
		} else {
			return WaitingStatus // 正在发送中
		}
	}
	return nil
}

// Send 发送短信
// signName = 短信签名名称
// phoneNumbers = 接收短信的手机号码
// templateCode = 短信模板Code
// templateParam = 短信模板变量对应的实际值
func (c *Aliyun) Send(signName string, phoneNumbers string, templateCode string, templateParam string) (bizID string, err error) {

	// 参数
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(signName),
		PhoneNumbers:  tea.String(phoneNumbers),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(templateParam),
	}
	response, err := c.client.SendSms(sendSmsRequest)
	if err != nil {
		return bizID, err
	}

	if !tea.BoolValue(util.EqualString(response.Body.Code, tea.String("OK"))) {
		return bizID, fmt.Errorf(tea.StringValue(response.Body.Message))
	}

	bizID = tea.StringValue(response.Body.BizId)
	if bizID == "" {
		return bizID, fmt.Errorf(tea.StringValue(response.Body.Message))
	}
	return bizID, nil
}
