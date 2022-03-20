package gossl

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// https://mritd.com/2021/05/31/golang-check-certificate-expiration-time/
func checkSSl(beforeTime time.Duration) error {
	c := &http.Client{
		Transport: &http.Transport{
			// 注意如果证书已过期，那么只有在关闭证书校验的情况下链接才能建立成功
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		// 10s 超时后认为服务挂了
		Timeout: 10 * time.Second,
	}
	resp, err := c.Get("https://mritd.com")
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	// 遍历所有证书
	for _, cert := range resp.TLS.PeerCertificates {
		// 检测证书是否已经过期
		if !cert.NotAfter.After(time.Now()) {
			return errors.New(fmt.Sprintf("Website [https://mritd.com] certificate has expired: %s", cert.NotAfter.Local().Format("2006-01-02 15:04:05")))
		}

		// 检测证书距离当前时间 是否小于 beforeTime
		// 例如 beforeTime = 7d，那么在证书过期前 6d 开始就发出警告
		if cert.NotAfter.Sub(time.Now()) < beforeTime {
			return errors.New(fmt.Sprintf("Website [https://mritd.com] certificate will expire, remaining time: %fh", cert.NotAfter.Sub(time.Now()).Hours()))
		}
	}
	return nil
}
