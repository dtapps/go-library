package gostorage

type Options struct {
	debug bool // 是否开启调试

	accessKeyId     string // 阿里云/百度云
	accessKeySecret string // 阿里云
	region          string // 地域节点
	bucket          string // 存储空间名称
}

type Option struct {
	F func(o *Options)
}

func NewOptions(opts []Option) *Options {
	options := &Options{}
	options.Apply(opts)
	return options
}

func (o *Options) Apply(opts []Option) {
	for _, op := range opts {
		op.F(o)
	}
}

// 设置 debug
func WithDebug(debug bool) Option {
	return Option{F: func(o *Options) {
		o.debug = debug
	}}
}

// 设置 accessKeyId
func WithAccessKeyId(accessKeyId string) Option {
	return Option{F: func(o *Options) {
		o.accessKeyId = accessKeyId
	}}
}

// 设置 accessKeySecret
func WithAccessKeySecret(accessKeySecret string) Option {
	return Option{F: func(o *Options) {
		o.accessKeySecret = accessKeySecret
	}}
}

// 设置 region
func WithRegion(region string) Option {
	return Option{F: func(o *Options) {
		o.region = region
	}}
}

// 设置 bucket
func WithBucket(bucket string) Option {
	return Option{F: func(o *Options) {
		o.bucket = bucket
	}}
}
