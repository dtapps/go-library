package wechatpayapiv2

import "crypto/tls"

func GetP12ToPem(cert, key string) (*tls.Certificate, error) {
	pemCert, err := tls.X509KeyPair([]byte(cert), []byte(key))
	return &pemCert, err
}
