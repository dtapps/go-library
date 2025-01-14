package gorequest

import (
	"errors"
	"strings"
)

// https://developers.weixin.qq.com/miniprogram/dev/framework/ability/network.html#%E4%BD%BF%E7%94%A8%E9%99%90%E5%88%B6
func isWechatMiniProgramRequestReferer(userAgent, referer string, appid string) error {
	// 判断结尾
	suffix := strings.HasSuffix(referer, "/page-frame.html")
	if suffix {
		// 判断开头
		prefix := strings.HasPrefix(referer, "https://servicewechat.com/")
		if prefix {
			// 判断加上appid
			prefixAppid := strings.HasPrefix(referer, "https://servicewechat.com/"+appid)
			if prefixAppid {
				return isWechatMiniProgramRequestUserAgent(userAgent)
			}
			return isWechatMiniProgramRequestUserAgent(userAgent)
		}
		return errors.New("格式固定不对2")
	}
	return errors.New("格式固定不对1")
}

func isWechatMiniProgramRequestUserAgent(userAgent string) error {
	if strings.Contains(userAgent, "MicroMessenger") {
		return nil
	}
	return errors.New("伪装数据")
}
