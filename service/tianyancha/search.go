package tianyancha

import (
	"fmt"
	"net/http"
)

type SearchHumanSuggestResult struct {
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

func (app *App) SearchHumanSuggest(key string) (body []byte, err error) {
	body, err = app.request(fmt.Sprintf("https://www.tianyancha.com/search/humanSuggest.json?key=%s", key), map[string]interface{}{}, http.MethodGet)
	return body, err
}