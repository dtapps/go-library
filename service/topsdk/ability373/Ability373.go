package ability373

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability373/request"
	"topsdk/ability373/response"
	"topsdk/util"
)

type Ability373 struct {
	Client *topsdk.TopClient
}

func NewAbility373(client *topsdk.TopClient) *Ability373 {
	return &Ability373{client}
}

/*
   聚划算商品搜索接口
*/
func (ability *Ability373) TaobaoJuItemsSearch(req *request.TaobaoJuItemsSearchRequest) (*response.TaobaoJuItemsSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability373 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.ju.items.search", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJuItemsSearchResponse{}
	if err != nil {
		log.Fatal("taobaoJuItemsSearch error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
