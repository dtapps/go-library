package wechatpayapiv3

func (c *Client) SignDecrypt(aesKey, associatedData, nonce, ciphertext string) ([]byte, error) {
	return c.decryptGCM(aesKey, nonce, ciphertext, associatedData)
}
