package ability382

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability382/request"
	"topsdk/ability382/response"
	"topsdk/util"
)

type Ability382 struct {
	Client *topsdk.TopClient
}

func NewAbility382(client *topsdk.TopClient) *Ability382 {
	return &Ability382{client}
}

/*
   淘宝客-推广者-红包领取状态查询
*/
func (ability *Ability382) TaobaoTbkDgVegasSendStatus(req *request.TaobaoTbkDgVegasSendStatusRequest) (*response.TaobaoTbkDgVegasSendStatusResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability382 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.send.status", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasSendStatusResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasSendStatus error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
