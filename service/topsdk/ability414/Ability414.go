package ability414

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability414/request"
	"topsdk/ability414/response"
	"topsdk/util"
)

type Ability414 struct {
	Client *topsdk.TopClient
}

func NewAbility414(client *topsdk.TopClient) *Ability414 {
	return &Ability414{client}
}

/*
   淘宝客-推广者-维权退款订单查询
*/
func (ability *Ability414) TaobaoTbkRelationRefund(req *request.TaobaoTbkRelationRefundRequest) (*response.TaobaoTbkRelationRefundResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability414 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.relation.refund", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkRelationRefundResponse{}
	if err != nil {
		log.Fatal("taobaoTbkRelationRefund error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-推广者-所有订单查询
*/
func (ability *Ability414) TaobaoTbkOrderDetailsGet(req *request.TaobaoTbkOrderDetailsGetRequest) (*response.TaobaoTbkOrderDetailsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability414 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.order.details.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkOrderDetailsGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkOrderDetailsGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
