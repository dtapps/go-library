package gojobs

import "log"

func (c *Client) Println(isPrint bool, v ...any) {
	if isPrint {
		log.Println(v)
	}
}
