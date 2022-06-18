package gojobs

import "xorm.io/xorm"

// Xorm数据库驱动
type jobsXorm struct {
	db *xorm.Engine
}

// NewJobsXorm 初始化
func NewJobsXorm(db *xorm.Engine) *jobsXorm {
	var (
		j = &jobsXorm{}
	)
	j.db = db
	return j
}

func (j *jobsXorm) Run() {

}
