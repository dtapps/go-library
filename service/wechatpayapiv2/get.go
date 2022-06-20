package wechatpayapiv2

import "crypto/tls"

func (c *Client) GetAppId() string {
	return c.config.AppId
}

func (c *Client) GetMchId() string {
	return c.config.MchId
}

func (c *Client) GetMchKey() string {
	return c.config.MchKey
}

func (c *Client) GetCertString() string {
	return c.config.CertString
}

func (c *Client) GetKeyString() string {
	return c.config.KeyString
}

func (c *Client) P12ToPem() (*tls.Certificate, error) {
	pemCert, err := tls.X509KeyPair([]byte(c.GetCertString()), []byte(c.GetKeyString()))
	return &pemCert, err
}
