package defaultability

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/defaultability/request"
	"topsdk/defaultability/response"
	"topsdk/util"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
   淘宝客-推广者-淘口令预警及拦截查询
*/
func (ability *Defaultability) TaobaoTbkDgTpwdRiskReport(req *request.TaobaoTbkDgTpwdRiskReportRequest) (*response.TaobaoTbkDgTpwdRiskReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.tpwd.risk.report", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgTpwdRiskReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgTpwdRiskReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-服务商-淘口令预警及拦截查询
*/
func (ability *Defaultability) TaobaoTbkScTpwdRiskReport(req *request.TaobaoTbkScTpwdRiskReportRequest, session string) (*response.TaobaoTbkScTpwdRiskReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.tpwd.risk.report", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScTpwdRiskReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScTpwdRiskReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-裂变淘礼金创建
*/
func (ability *Defaultability) TaobaoTbkDgVegasLbtljCreate(req *request.TaobaoTbkDgVegasLbtljCreateRequest) (*response.TaobaoTbkDgVegasLbtljCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.lbtlj.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasLbtljCreateResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasLbtljCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
