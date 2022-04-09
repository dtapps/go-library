package ability1826

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability1826/request"
	"topsdk/ability1826/response"
	"topsdk/util"
)

type Ability1826 struct {
	Client *topsdk.TopClient
}

func NewAbility1826(client *topsdk.TopClient) *Ability1826 {
	return &Ability1826{client}
}

/*
   淘宝客-推广者-淘礼金发放及使用报表(将停用，按描述说明换新API)
*/
func (ability *Ability1826) TaobaoTbkDgVegasTljInstanceReport(req *request.TaobaoTbkDgVegasTljInstanceReportRequest) (*response.TaobaoTbkDgVegasTljInstanceReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1826 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.tlj.instance.report", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasTljInstanceReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasTljInstanceReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-淘礼金暂停发放
*/
func (ability *Ability1826) TaobaoTbkDgVegasTljStop(req *request.TaobaoTbkDgVegasTljStopRequest) (*response.TaobaoTbkDgVegasTljStopResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1826 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.tlj.stop", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasTljStopResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasTljStop error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-淘礼金创建
*/
func (ability *Ability1826) TaobaoTbkDgVegasTljCreate(req *request.TaobaoTbkDgVegasTljCreateRequest) (*response.TaobaoTbkDgVegasTljCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1826 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.tlj.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasTljCreateResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasTljCreate error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-淘礼金效果数据
*/
func (ability *Ability1826) TaobaoTbkDgVegasTljReport(req *request.TaobaoTbkDgVegasTljReportRequest) (*response.TaobaoTbkDgVegasTljReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability1826 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.tlj.report", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasTljReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasTljReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
