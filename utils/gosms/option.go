package gosms

type Options struct {
	accessKeyId     string // 阿里云
	accessKeySecret string // 阿里云
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
