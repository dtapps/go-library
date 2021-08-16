package v2

import (
	"encoding/json"
)

type GetScheduleList struct {
	CinemaId int `json:"cinemaId"` // 影院id
}

type GetScheduleListResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			PlanType        string `json:"planType"`        // 影厅类型 2D 3D
			ShowTime        string `json:"showTime"`        // 放映时间
			NetPrice        int    `json:"netPrice"`        // 参考价，单位：分
			Language        string `json:"language"`        // 语言
			ShowDate        string `json:"showDate"`        //
			Duration        int    `json:"duration"`        // 时长,分钟
			ShowId          string `json:"showId"`          // 场次标识
			StopSellTime    string `json:"stopSellTime"`    // 停售时间
			CinemaId        int    `json:"cinemaId"`        // 影院id
			CinemaName      string `json:"cinemaName"`      //
			FilmId          int    `json:"filmId"`          // 影片id
			ScheduleArea    string `json:"scheduleArea"`    // 该排期的分区座位价格信息，当此字段有值的时候，代表座位里面支持分区价格。 如果调用的是秒出票下单， 那价格必须计算正确，才能正确出票成功，即必须处理好座位分区价格
			FilmName        string `json:"filmName"`        // 影片名字
			HallName        string `json:"hallName"`        // 影厅名
			ShowVersionType string `json:"showVersionType"` // 场次类型
		} `json:"list"`
		DiscountRule struct {
			UpDiscountRate   float64 `json:"upDiscountRate"`   // 影院最高成本折扣，当价格大于等于39元时候，可取此字段
			DownDiscountRate float64 `json:"downDiscountRate"` // 影院最高成本折扣，当价格小于39元时候，可取此字段
		} `json:"discountRule"`
	} `json:"data"`
	Success bool `json:"success"`
}

// GetScheduleList 场次排期 https://www.showdoc.com.cn/1154868044931571/5866708808899217
func (app *App) GetScheduleList(param GetScheduleList) (result GetScheduleListResult, err error) {
	// api params
	params := map[string]interface{}{}
	b, _ := json.Marshal(&param)
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	for k, v := range m {
		params[k] = v
	}
	body, err := app.request("movieapi/movie-info/get-schedule-list", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
