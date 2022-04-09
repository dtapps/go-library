package pintoto

import (
	"encoding/json"
)

type GetSeat struct {
	ShowId string `json:"showId"` // 场次标识
}

type GetSeatResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		SeatData struct {
			Restrictions int            `json:"restrictions"`
			Seats        []GetSeatSeats `json:"seats"`
		} `json:"seatData"`
	} `json:"data"`
	Success bool `json:"success"`
}

type GetSeatSeats struct {
	Area       string `json:"area"`       // 本座位所在的区域，根据场次排期接口的 scheduleArea 字段， 可得到当前座位的分区价格
	ColumnNo   string `json:"columnNo"`   // 列
	Lovestatus int    `json:"lovestatus"` // 0为非情侣座；1为情侣座左；2为情侣座右
	RowNo      string `json:"rowNo"`      // 行
	SeatId     string `json:"seatId"`     // 座位标识符，锁座位和秒出票的时候需要用到
	SeatNo     string `json:"seatNo"`     // 座位名
	Status     string `json:"status"`     // N可售，LK不可售
}

// GetSeat 座位 https://www.showdoc.com.cn/1154868044931571/5866824368760475
func (app *App) GetSeat(showId string) (result GetSeatResult, err error) {
	// 参数
	param := NewParams()
	param.Set("showId", showId)
	// 转换
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://movieapi2.pintoto.cn/movieapi/movie-info/get-seat", params)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
