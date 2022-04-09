package ability3280

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability3280/request"
	"topsdk/ability3280/response"
	"topsdk/util"
)

type Ability3280 struct {
	Client *topsdk.TopClient
}

func NewAbility3280(client *topsdk.TopClient) *Ability3280 {
	return &Ability3280{client}
}

/*
   淘宝客-推广者-淘口令回流数据查询
*/
func (ability *Ability3280) TaobaoTbkDgTpwdReportGet(req *request.TaobaoTbkDgTpwdReportGetRequest) (*response.TaobaoTbkDgTpwdReportGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability3280 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tbk.dg.tpwd.report.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTbkDgTpwdReportGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkDgTpwdReportGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
