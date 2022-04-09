package ability2474

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability2474/request"
	"topsdk/ability2474/response"
	"topsdk/util"
)

type Ability2474 struct {
	Client *topsdk.TopClient
}

func NewAbility2474(client *topsdk.TopClient) *Ability2474 {
	return &Ability2474{client}
}

/*
   淘宝客-推广者-查询红包发放个数
*/
func (ability *Ability2474) TaobaoTbkDgVegasSendReport(req *request.TaobaoTbkDgVegasSendReportRequest) (*response.TaobaoTbkDgVegasSendReportResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability2474 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.vegas.send.report", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgVegasSendReportResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgVegasSendReport error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
