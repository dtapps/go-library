package ability3261

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability3261/request"
	"topsdk/ability3261/response"
	"topsdk/util"
)

type Ability3261 struct {
	Client *topsdk.TopClient
}

func NewAbility3261(client *topsdk.TopClient) *Ability3261 {
	return &Ability3261{client}
}

/*
   淘宝客-推广者-CPA活动执行明细
*/
func (ability *Ability3261) TaobaoTbkDgCpaActivityDetail(req *request.TaobaoTbkDgCpaActivityDetailRequest) (*response.TaobaoTbkDgCpaActivityDetailResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability3261 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.cpa.activity.detail", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgCpaActivityDetailResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgCpaActivityDetail error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-任务奖励效果报表
*/
func (ability *Ability3261) TaobaoTbkDgCpaActivityReport(req *request.TaobaoTbkDgCpaActivityReportRequest) (*response.TaobaoTbkDgCpaActivityReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability3261 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.cpa.activity.report", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgCpaActivityReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgCpaActivityReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
