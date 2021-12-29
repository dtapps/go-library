package dingdanxia

import (
	"gopkg.in/dtapps/go-library.v3/utils/gotime"
	"log"
	"testing"
)

func TestApp_JdJyOrderDetails(t *testing.T) {
	param := NewParams()
	param.Set("startTime", gotime.Current().BeforeDay(28).Timestamp())
	param.Set("endTime", gotime.Current().Timestamp())
	param.Set("itemsPerPage", 60)
	param.Set("curPage", 1)
	param.Set("type", 1)
	log.Printf("%+v\n", app.JdJyOrderDetails(param))
}
