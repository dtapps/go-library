package dingdanxia

import (
	"gopkg.in/dtapps/go-library.v3/utils/gotime"
	"log"
	"testing"
)

func TestApp_WaimaiMeituanOrders(t *testing.T) {
	param := NewParams()
	param.Set("start_time", gotime.Current().BeforeDay(28).Timestamp())
	param.Set("end_time", gotime.Current().Timestamp())
	log.Printf("%+v\n", app.WaimaiMeituanOrders(param))
}
