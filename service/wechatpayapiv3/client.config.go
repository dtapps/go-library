package wechatpayapiv3

import (
	"crypto/rsa"
	"crypto/x509"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetMchId() string {
	return c.config.mchId
}

func (c *Client) GetApiV3() string {
	return c.config.apiV3
}

func (c *Client) GetCertificateSerialNo() string {
	return c.config.certificateSerialNo
}

func (c *Client) GetCertificate() *x509.Certificate {
	return c.config.certificate
}

func (c *Client) GetPrivateKey() *rsa.PrivateKey {
	return c.config.privateKey
}

func (c *Client) GetPublicKeyID() string {
	return c.config.publicKeyID
}

func (c *Client) GetPublicKey() *rsa.PublicKey {
	return c.config.publicKey
}
