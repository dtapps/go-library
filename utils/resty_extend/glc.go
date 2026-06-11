package resty_extend

import (
	"fmt"

	"github.com/gotoeasy/glang/cmn"
)

type GlcLogger struct{}

func (l *GlcLogger) Errorf(format string, v ...any) {
	cmn.Error(fmt.Sprintf(format, v...))
}

func (l *GlcLogger) Warnf(format string, v ...any) {
	cmn.Warn(fmt.Sprintf(format, v...))
}

func (l *GlcLogger) Debugf(format string, v ...any) {
	cmn.Debug(fmt.Sprintf(format, v...))
}
