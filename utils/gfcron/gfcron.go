package gfcron

import (
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
)

// https://goframe.org/pages/viewpage.action?pageId=30736411

var (
	Cron *gcron.Cron
	Ctx  = gctx.New()
)
