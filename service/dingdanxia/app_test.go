package dingdanxia

import (
	"log"
	"testing"

	"github.com/dtapps/go-library/utils/gotime"
)

var app = App{
	ApiKey: "ZTe5tqKhIc6nm8HkJkqj5hmpmAjs1WTy",
}

func TestApp(t *testing.T) {
	param := NewParams()
	param.Set("start_time", gotime.Current().BeforeDay(28).Timestamp())
	param.Set("end_time", gotime.Current().Timestamp())
	res, err := app.WaimaiMeituanOrders(param)
	log.Println(res)
	log.Printf("%s\n", res)
	log.Printf("%v\n", res)
	log.Println(err)
}
