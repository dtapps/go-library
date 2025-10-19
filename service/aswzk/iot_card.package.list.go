package aswzk

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type IotCardPackageListResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data []struct {
		PackageCode           string  `json:"package_code"`             // 套餐编号
		PackageName           string  `json:"package_name"`             // 套餐名称
		PackageFlow           int     `json:"package_flow"`             // 套餐流量 单位：KB
		PackageMonthLimitFlow int     `json:"package_month_limit_flow"` // 套餐月限制流量 0=不限制 单位：KB
		PackageDay            int     `json:"package_day"`              // 套餐天数
		PackageType           string  `json:"package_type"`             // 套餐类型 day=天 month=月 half_year=半年 year=年
		PackageOperator       string  `json:"package_operator"`         // 套餐运营商 mobile=移动 unicom=联通 telecom=电信
		PriceSelling          float64 `json:"price_selling"`            // 销售价格
	} `json:"data,omitempty"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

// IotCardPackageList 物联卡套餐列表
func (c *Client) IotCardPackageList(ctx context.Context, iccid string, notMustParams ...*gorequest.Params) (response IotCardPackageListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("iccid", iccid)

	// 请求
	err = c.request(ctx, "iot_card/package/list", params, http.MethodGet, &response)
	return
}
