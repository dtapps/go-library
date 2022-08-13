package gostring

import (
	"crypto/rand"
	"fmt"
	"time"
)

// GetUuId 由 32 个十六进制数字组成，以 6 个组显示，由连字符 - 分隔
func GetUuId() string {
	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	numRead, err := rand.Read(buff)
	if numRead != len(buff) || err != nil {
		return ""
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}
