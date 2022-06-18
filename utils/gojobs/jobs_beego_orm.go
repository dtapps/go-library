package gojobs

import "github.com/beego/beego/v2/client/orm"

type JobsBeegoOrm struct {
	Db *orm.Ormer
}

func NewJobsBeegoOrm(db *orm.Ormer) *JobsBeegoOrm {
	var (
		jobsBeegoOrm = &JobsBeegoOrm{}
	)
	jobsBeegoOrm.Db = db
	return jobsBeegoOrm
}
