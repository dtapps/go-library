package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func (c *Client) sign(t int64) string {
	secStr := fmt.Sprintf("%d\n%s", t, c.GetSecret())
	hmac256 := hmac.New(sha256.New, []byte(c.GetSecret()))
	hmac256.Write([]byte(secStr))
	result := hmac256.Sum(nil)
	return base64.StdEncoding.EncodeToString(result)
}
