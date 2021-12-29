package wechatminiprogram

import (
	"testing"
)

func TestApp_GetCallBackIp(t *testing.T) {
	t.Log(app)
	result := app.GetCallBackIp()
	t.Logf("Response：%s", result.GetCallBackIpResponse)
	t.Logf("Err：%s", result.Err)
	t.Logf("Byte：%v", result.Byte)
}
