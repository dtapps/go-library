package ability2138

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability2138/request"
	"topsdk/ability2138/response"
	"topsdk/util"
)

type Ability2138 struct {
	Client *topsdk.TopClient
}

func NewAbility2138(client *topsdk.TopClient) *Ability2138 {
	return &Ability2138{client}
}

/*
   淘宝客-推广者-新用户订单明细查询
*/
func (ability *Ability2138) TaobaoTbkDgNewuserOrderGet(req *request.TaobaoTbkDgNewuserOrderGetRequest) (*response.TaobaoTbkDgNewuserOrderGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2138 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.newuser.order.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgNewuserOrderGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgNewuserOrderGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
