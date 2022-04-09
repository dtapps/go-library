package ability425

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability425/request"
	"topsdk/ability425/response"
	"topsdk/util"
)

type Ability425 struct {
	Client *topsdk.TopClient
}

func NewAbility425(client *topsdk.TopClient) *Ability425 {
	return &Ability425{client}
}

/*
   淘宝客-公用-私域用户邀请码生成
*/
func (ability *Ability425) TaobaoTbkScInvitecodeGet(req *request.TaobaoTbkScInvitecodeGetRequest, session string) (*response.TaobaoTbkScInvitecodeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability425 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.invitecode.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScInvitecodeGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScInvitecodeGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-公用-私域用户备案
*/
func (ability *Ability425) TaobaoTbkScPublisherInfoSave(req *request.TaobaoTbkScPublisherInfoSaveRequest, session string) (*response.TaobaoTbkScPublisherInfoSaveResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability425 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.publisher.info.save", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScPublisherInfoSaveResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScPublisherInfoSave error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   淘宝客-公用-私域用户备案信息查询
*/
func (ability *Ability425) TaobaoTbkScPublisherInfoGet(req *request.TaobaoTbkScPublisherInfoGetRequest, session string) (*response.TaobaoTbkScPublisherInfoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability425 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tbk.sc.publisher.info.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTbkScPublisherInfoGetResponse{}
	if err != nil {
		log.Fatal("taobaoTbkScPublisherInfoGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
