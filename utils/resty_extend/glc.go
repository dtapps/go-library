package resty_extend

import (
	"fmt"

	"github.com/gotoeasy/glang/cmn"
)

type GlcLogger struct{}

func (l *GlcLogger) Errorf(format string, v ...interface{}) {
	cmn.Error(fmt.Sprintf(format, v...))
}

func (l *GlcLogger) Warnf(format string, v ...interface{}) {
	cmn.Warn(fmt.Sprintf(format, v...))
}

func (l *GlcLogger) Debugf(format string, v ...interface{}) {
	cmn.Debug(fmt.Sprintf(format, v...))
}
