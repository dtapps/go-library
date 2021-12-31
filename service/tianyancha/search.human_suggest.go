package tianyancha

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SearchHumanSuggestResponse struct {
	Data struct {
		Id          int         `json:"id"`
		HumanName   interface{} `json:"humanName"`
		DistinctNum int         `json:"distinctNum"`
		ViewNum     int         `json:"viewNum"`
		ResultCount int         `json:"resultCount"`
		ResultList  []struct {
			Name            string        `json:"name"`
			Hid             int64         `json:"hid"`
			HeadUrl         interface{}   `json:"headUrl"`
			Introduction    interface{}   `json:"introduction"`
			Event           interface{}   `json:"event"`
			BossCertificate int           `json:"bossCertificate"`
			CompanyNum      int           `json:"companyNum"`
			Office          []interface{} `json:"office"`
			Companys        interface{}   `json:"companys"`
			PartnerNum      int           `json:"partnerNum"`
			CoopCount       int           `json:"coopCount"`
			Partners        interface{}   `json:"partners"`
			Cid             int64         `json:"cid"`
			TypeJoin        interface{}   `json:"typeJoin"`
			Alias           interface{}   `json:"alias"`
			ServiceType     int           `json:"serviceType"`
			ServiceCount    int           `json:"serviceCount"`
			OfficeV1        []interface{} `json:"officeV1"`
			Pid             interface{}   `json:"pid"`
			Role            interface{}   `json:"role"`
		} `json:"resultList"`
		TotalPage   int         `json:"totalPage"`
		CurrentPage int         `json:"currentPage"`
		RealName    interface{} `json:"realName"`
		AdviceQuery interface{} `json:"adviceQuery"`
	} `json:"data"`
	VipMessage string `json:"vipMessage"`
	Special    string `json:"special"`
	State      string `json:"state"`
}

type SearchHumanSuggestResult struct {
	Result SearchHumanSuggestResponse // 结果
	Body   []byte                     // 内容
	Err    error                      // 错误
}

func NewSearchHumanSuggestResult(result SearchHumanSuggestResponse, body []byte, err error) *SearchHumanSuggestResult {
	return &SearchHumanSuggestResult{Result: result, Body: body, Err: err}
}

func (app *App) SearchHumanSuggest(key string) *SearchHumanSuggestResult {
	body, err := app.request(fmt.Sprintf("https://www.tianyancha.com/search/humanSuggest.json?key=%s", key), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response SearchHumanSuggestResponse
	err = json.Unmarshal(body, &response)
	return NewSearchHumanSuggestResult(response, body, err)
}
