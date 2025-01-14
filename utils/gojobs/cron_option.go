package gojobs

type CronOption func(*Cron)

// WithCronLog 日志
func WithCronLog() CronOption {
	return func(c *Cron) {
		c.option.log = true
	}
}
