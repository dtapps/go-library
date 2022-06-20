package wechatoffice

import "errors"

func (c *Client) pkcs7Unpaid(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("invalid block size")
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, errors.New("invalid PKCS7 data")
	}
	d := data[len(data)-1]
	n := int(d)
	if n == 0 || n > len(data) {
		return nil, errors.New("invalid padding on input")
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != d {
			return nil, errors.New("invalid padding on input")
		}
	}
	return data[:len(data)-n], nil
}
