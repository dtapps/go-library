package gojobs

//
//type JobsOption func(*JobsCron)
//
//// WithRedis 缓存服务驱动
//func WithRedis(db *goredis.Client) JobsOption {
//	return func(opts *JobsCron) {
//		opts.redis = db
//	}
//}
//
//// WithGorm 数据库服务驱动
//func WithGorm(db *gorm.DB) JobsOption {
//	return func(opts *JobsCron) {
//		opts.db = db
//	}
//}
//
//// WithMainService 是否主要服务(主要服务可删除过期服务)
//func WithMainService(status int) JobsOption {
//	return func(opts *JobsCron) {
//		opts.mainService = status
//	}
//}
