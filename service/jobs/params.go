package jobs

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ParamsOrderType = "order"

// ParamsOrderId 订单任务
type ParamsOrderId struct {
	OrderId string `json:"order_id,omitempty"`
}

var ParamsMerchantGoldenBeanType = "merchant.golden_bean"

// ParamsMerchantUserIdOpenid 商家金豆任务
type ParamsMerchantUserIdOpenid struct {
	MerchantUserId int64  `json:"merchant_user_id,omitempty"`
	Openid         string `json:"openid,omitempty"`
}

var ParamsNewServiceType = "new_service"

// ParamsTaskId 企业自定义任务
type ParamsTaskId struct {
	TaskId int64 `json:"task_id,omitempty"`
}

var ParamsNewServiceNextType = "new_service.next"

// ParamsTaskIdNext 企业自定义下一步任务
type ParamsTaskIdNext struct {
	TaskId         int64 `json:"task_id,omitempty"`
	MerchantUserId int64 `json:"merchant_user_id,omitempty"`
	CurrentNumber  int   `json:"current_number,omitempty"`
	MaxNumber      int   `json:"max_number,omitempty"`
}

var ParamsWechatType = "wechat"

// ParamsWechat 微信任务
type ParamsWechat struct {
	Appid string `json:"appid,omitempty"`
	Type  string `json:"type,omitempty"`
}

var ParamsTeamInvType = "team.inv"

// ParamsTeamInv 团队邀请任务
type ParamsTeamInv struct {
	MerchantUserId int64  `json:"merchant_user_id,omitempty"`
	Openid         string `json:"openid,omitempty"`
	ShareOpenid    string `json:"share_openid,omitempty"`
}

var ParamsRepairMerchantAccountQuantityLevelType = "repair.merchant.account.quantity.level"

// ParamsRepairMerchantAccountQuantityLevel 修复商家账号数量下一步任务
type ParamsRepairMerchantAccountQuantityLevel struct {
	Level int `json:"level,omitempty"`
}

var ParamsKashangwlType = "kashangwl"

type ParamsKashangwlId struct {
	ProductID int64 `json:"product_id"`
}

// Params 任务参数
func Params(v Task) (response CronParamsResp, err error) {
	switch v.ParamsType {
	case ParamsOrderType:
		// 订单任务
		var resp ParamsOrderId
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.OrderId == "" {
			return response, errors.New("参数不存在")
		}
		response.ParamsOrderId.OrderId = resp.OrderId
		return response, nil
	case ParamsMerchantGoldenBeanType:
		// 商家金豆任务
		var resp ParamsMerchantUserIdOpenid
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.MerchantUserId == 0 || resp.Openid == "" {
			return response, errors.New("参数不存在")
		}
		response.ParamsMerchantUserIdOpenid.MerchantUserId = resp.MerchantUserId
		response.ParamsMerchantUserIdOpenid.Openid = resp.Openid
		return response, nil
	case ParamsNewServiceType:
		// 企业自定义任务
		var resp ParamsTaskId
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.TaskId == 0 {
			return response, errors.New("参数不存在")
		}
		response.ParamsTaskId.TaskId = resp.TaskId
		return response, nil
	case ParamsNewServiceNextType:
		// 企业自定义下一步任务
		var resp ParamsTaskIdNext
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.TaskId == 0 || resp.MerchantUserId == 0 || resp.CurrentNumber == 0 || resp.MaxNumber == 0 {
			return response, errors.New("参数不存在")
		}
		response.ParamsTaskIdNext.TaskId = resp.TaskId
		response.ParamsTaskIdNext.MerchantUserId = resp.MerchantUserId
		response.ParamsTaskIdNext.CurrentNumber = resp.CurrentNumber
		response.ParamsTaskIdNext.MaxNumber = resp.MaxNumber
		return response, nil
	case ParamsWechatType:
		// 微信任务
		var resp ParamsWechat
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.Appid == "" || resp.Type == "" {
			return response, errors.New("参数不存在")
		}
		response.ParamsWechat.Appid = resp.Appid
		response.ParamsWechat.Type = resp.Type
		return response, nil
	case ParamsTeamInvType:
		// 团队邀请任务
		var resp ParamsTeamInv
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.MerchantUserId == 0 || resp.Openid == "" || resp.ShareOpenid == "" {
			return response, errors.New("参数不存在")
		}
		response.ParamsTeamInv.MerchantUserId = resp.MerchantUserId
		response.ParamsTeamInv.Openid = resp.Openid
		response.ParamsTeamInv.ShareOpenid = resp.ShareOpenid
		return response, nil
	case ParamsRepairMerchantAccountQuantityLevelType:
		// 修复商家账号数量下一步任务
		var resp ParamsRepairMerchantAccountQuantityLevel
		err = json.Unmarshal([]byte(v.Params), &resp)
		if err != nil {
			return response, errors.New(fmt.Sprintf("解析失败，%v", err))
		}
		if resp.Level == 0 {
			return response, errors.New("参数不存在")
		}
		response.ParamsRepairMerchantAccountQuantityLevel.Level = resp.Level
		return response, nil
	default:
		return response, nil
	}
}
