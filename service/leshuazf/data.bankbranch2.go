package leshuazf

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DataBankBranch2Response struct {
	RespCode    string `json:"respCode"`
	RespMsg     string `json:"respMsg"`
	ReqSerialNo string `json:"reqSerialNo"`
	Data        struct {
		Total    int `json:"total"`
		Page     int `json:"page"`
		PageSize int `json:"pageSize"`
		List     []struct {
			UnionpayCode       string      `json:"unionpayCode"`
			FinInstitutionCode string      `json:"finInstitutionCode"`
			BankArea           string      `json:"bankArea"`
			BankCity           string      `json:"bankCity"`
			BranchName         string      `json:"branchName"`
			CftAreaCode        string      `json:"cftAreaCode"`
			CftCityCode        string      `json:"cftCityCode"`
			UnionAreaCode      interface{} `json:"unionAreaCode"`
			UnionCityCode      interface{} `json:"unionCityCode"`
			FinInstitutionName string      `json:"finInstitutionName"`
			BankName           string      `json:"bankName"`
		} `json:"list"`
	} `json:"data"`
}

type DataBankBranch2Result struct {
	Result DataBankBranch2Response // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func newDataBankBranch2Result(result DataBankBranch2Response, body []byte, http gorequest.Response, err error) *DataBankBranch2Result {
	return &DataBankBranch2Result{Result: result, Body: body, Http: http, Err: err}
}

// DataBankBranch2 代理商通过联行号来查支行信息
// https://www.yuque.com/leshuazf/doc/dbmxyi#QYl0c
func (c *Client) DataBankBranch2(ctx context.Context, notMustParams ...gorequest.Params) *DataBankBranch2Result {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/data/bankbranch2", params, http.MethodPost)
	// 定义
	var response DataBankBranch2Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDataBankBranch2Result(response, request.ResponseBody, request, err)
}
