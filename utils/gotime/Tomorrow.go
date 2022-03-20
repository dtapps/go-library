package gotime

import "time"

// Tomorrow 明天
func Tomorrow() Pro {
	p := NewPro()
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		p.Time = time.Now().Add(time.Hour*8).AddDate(0, 0, +1)
	} else {
		p.Time = time.Now().In(location).AddDate(0, 0, +1)
	}
	return p
}
