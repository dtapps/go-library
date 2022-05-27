package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func (app *App) sign(t int64) string {
	secStr := fmt.Sprintf("%d\n%s", t, app.secret)
	hmac256 := hmac.New(sha256.New, []byte(app.secret))
	hmac256.Write([]byte(secStr))
	result := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}
