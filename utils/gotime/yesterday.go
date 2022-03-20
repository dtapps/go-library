package gotime

import (
	"time"
)

// Yesterday 昨天
func Yesterday() Pro {
	p := NewPro()
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		p.Time = time.Now().Add(time.Hour*8).AddDate(0, 0, -1)
	} else {
		p.Time = time.Now().In(location).AddDate(0, 0, -1)
	}
	return p
}
