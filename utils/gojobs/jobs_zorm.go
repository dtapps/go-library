package gojobs

import "gitee.com/chunanyong/zorm"

// Zorm数据库驱动
type jobsZorm struct {
	db *zorm.DBDao
}

// NewJobsZorm 初始化
func NewJobsZorm(db *zorm.DBDao) *jobsZorm {
	var (
		j = &jobsZorm{}
	)
	j.db = db
	return j
}

func (j *jobsZorm) Run() {

}
