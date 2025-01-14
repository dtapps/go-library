package kuaidi100

import (
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
)

func (c *Client) getSign(param string) string {
	return gomd5.ToUpper(fmt.Sprintf("%s%s%s", param, c.GetKey(), c.GetCustomer()))
}
