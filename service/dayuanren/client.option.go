package dayuanren

import "resty.dev/v3"

type Options struct {
	httpClient *resty.Client
}

type Option struct {
	F func(o *Options)
}

func NewOptions(opts []Option) *Options {
	options := &Options{
		httpClient: nil,
	}
	options.Apply(opts)
	return options
}

func (o *Options) Apply(opts []Option) {
	for _, op := range opts {
		op.F(o)
	}
}

// WithResty 设置自定义的Resty
func WithResty(httpClient *resty.Client) Option {
	return Option{F: func(o *Options) {
		o.httpClient = httpClient
	}}
}
