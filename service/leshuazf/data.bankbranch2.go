package leshuazf

import (
	"encoding/json"
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
	Err    error                   // 错误
}

func NewDataBankBranch2Result(result DataBankBranch2Response, body []byte, err error) *DataBankBranch2Result {
	return &DataBankBranch2Result{Result: result, Body: body, Err: err}
}

// DataBankBranch2 代理商通过联行号来查支行信息
// https://www.yuque.com/leshuazf/doc/dbmxyi#QYl0c
func (app *App) DataBankBranch2(notMustParams ...Params) *DataBankBranch2Result {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("data/bankbranch2", params, http.MethodPost)
	// 定义
	var response DataBankBranch2Response
	err = json.Unmarshal(body, &response)
	return NewDataBankBranch2Result(response, body, err)
}
