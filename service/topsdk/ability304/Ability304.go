package ability304

import (
	"errors"
	"log"
	"topsdk"
	"topsdk/ability304/request"
	"topsdk/ability304/response"
	"topsdk/util"
)

type Ability304 struct {
	Client *topsdk.TopClient
}

func NewAbility304(client *topsdk.TopClient) *Ability304 {
	return &Ability304{client}
}

/*
   业务文件获取
*/
func (ability *Ability304) TaobaoFilesGet(req *request.TaobaoFilesGetRequest) (*response.TaobaoFilesGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.files.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoFilesGetResponse{}
	if err != nil {
		log.Fatal("taobaoFilesGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   刷新Access Token
*/
func (ability *Ability304) TaobaoTopAuthTokenRefresh(req *request.TaobaoTopAuthTokenRefreshRequest) (*response.TaobaoTopAuthTokenRefreshResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.top.auth.token.refresh", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTopAuthTokenRefreshResponse{}
	if err != nil {
		log.Fatal("taobaoTopAuthTokenRefresh error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   获取授权账号对应的OpenUid
*/
func (ability *Ability304) TaobaoOpenuidGet(req *request.TaobaoOpenuidGetRequest, session string) (*response.TaobaoOpenuidGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.openuid.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOpenuidGetResponse{}
	if err != nil {
		log.Fatal("taobaoOpenuidGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   通过订单获取对应买家的openUID
*/
func (ability *Ability304) TaobaoOpenuidGetBytrade(req *request.TaobaoOpenuidGetBytradeRequest, session string) (*response.TaobaoOpenuidGetBytradeResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.openuid.get.bytrade", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOpenuidGetBytradeResponse{}
	if err != nil {
		log.Fatal("taobaoOpenuidGetBytrade error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   获取TOP通道解密秘钥
*/
func (ability *Ability304) TaobaoTopSecretGet(req *request.TaobaoTopSecretGetRequest) (*response.TaobaoTopSecretGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.top.secret.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTopSecretGetResponse{}
	if err != nil {
		log.Fatal("taobaoTopSecretGet error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   通过mixnick转换openuid
*/
func (ability *Ability304) TaobaoOpenuidGetBymixnick(req *request.TaobaoOpenuidGetBymixnickRequest) (*response.TaobaoOpenuidGetBymixnickResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.openuid.get.bymixnick", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoOpenuidGetBymixnickResponse{}
	if err != nil {
		log.Fatal("taobaoOpenuidGetBymixnick error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
   sdk信息回调
*/
func (ability *Ability304) TaobaoTopSdkFeedbackUpload(req *request.TaobaoTopSdkFeedbackUploadRequest) (*response.TaobaoTopSdkFeedbackUploadResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Ability304 topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.top.sdk.feedback.upload", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTopSdkFeedbackUploadResponse{}
	if err != nil {
		log.Fatal("taobaoTopSdkFeedbackUpload error", err)
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
