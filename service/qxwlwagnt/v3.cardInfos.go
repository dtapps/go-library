package qxwlwagnt

import (
	"context"
	"net/http"
	"strings"

	"go.dtapp.net/library/utils/gorequest"
)

type V3CardInfosResponse struct {
	Iccid     string `json:"iccid"` // 物联网号码的ICCID
	IccidData struct {
		Iccid      string `json:"iccid"`      // 物联网号码的ICCID
		BelongId   string `json:"belongId"`   // 通道编号
		Msisdn     string `json:"msisdn"`     // 物联网号码的MSISDN
		Balance    string `json:"balance"`    // 卡余额（元）
		UsedGprs   string `json:"usedGprs"`   // 当月流量使用量（MB）
		ProdTotal  string `json:"prodTotal"`  // 套餐总量（MB）
		ProdName   string `json:"prodName"`   // 套餐名称
		SimStatus  string `json:"simStatus"`  // SIM卡状态
		GprsStatus string `json:"gprsStatus"` // 网络状态 Y正常；F断网
		OpenDate   string `json:"openDate"`   // 出库时间/开卡时间
		ActiveTime string `json:"activeTime"` // 激活时间
		ExpTime    string `json:"expTime"`    // 过期时间
		UpdateTime string `json:"updateTime"` // 数据同步时间
		Imsi       string `json:"imsi"`       // 物联网号码的imsi
		FlowPoolId string `json:"flowPoolId"` // 流量池编号
	} `json:"iccidData"` // 物联网号码的ICCID相关数据
}

// V3CardInfos 批量物联卡信息查询
// http://docs.konyun.net/web/#/71/2391
func (c *Client) V3CardInfos(ctx context.Context, iccids []string, notMustParams ...*gorequest.Params) (response CommonResponse[[]V3CardInfosResponse], err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("iccids", strings.Join(iccids, "_")) // 集成电路卡识别码，IC 卡的唯一识别号码(多个iccid之间_分割,最多不超过200个)

	// 请求
	err = c.Request(ctx, "/api/v3/cardInfos", params, http.MethodGet, &response)
	return
}
