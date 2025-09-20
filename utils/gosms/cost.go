package gosms

import "errors"

var (
	SuccessStatus = errors.New("发送成功")
	WaitingStatus = errors.New("等待回执")
	FailureStatus = errors.New("发送失败")
)
