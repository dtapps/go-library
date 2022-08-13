package gojobs

import (
	"context"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"gorm.io/gorm"
)

// RefreshIp 刷新Ip
func (j *JobsGorm) RefreshIp(ctx context.Context, tx *gorm.DB) {
	xip := goip.GetOutsideIp(ctx)
	if j.config.outsideIp == "" || j.config.outsideIp == "0.0.0.0" {
		return
	}
	if j.config.outsideIp == xip {
		return
	}
	tx.Where("ips = ?", j.config.outsideIp).Delete(&jobs_gorm_model.TaskIp{}) // 删除
	j.config.outsideIp = xip
}
