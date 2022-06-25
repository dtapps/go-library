package gojobs

import (
	"go.dtapp.net/library/utils/goip"
	"go.dtapp.net/library/utils/gojobs/jobs_gorm_model"
	"gorm.io/gorm"
)

// RefreshIp 刷新Ip
func (j *JobsGorm) RefreshIp(tx *gorm.DB) {
	xip := goip.GetOutsideIp()
	if j.config.OutsideIp == "" || j.config.OutsideIp == "0.0.0.0" {
		return
	}
	if j.config.OutsideIp == xip {
		return
	}
	tx.Where("ips = ?", j.config.OutsideIp).Delete(&jobs_gorm_model.TaskIp{}) // 删除
	j.config.OutsideIp = xip
}
