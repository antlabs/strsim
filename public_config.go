package strsim

type option struct {
	ignore int //
	ngram  int // dice coefficient 算法会需要用到
	cmp    func(s1, s2 string) float64
}

// 调用Option接口设置option
func (o *option) fillOption(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(o)
	}
}

type Option interface {
	Apply(*option)
}

type OptionFunc func(*option)

func (o OptionFunc) Apply(opt *option) {
	o(opt)
}

//忽略大小写
func IgnoreCase() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ignore |= ignoreCase
	})
}

//忽略空白字符
func IgnoreSpace() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ignore |= ignoreSpace
	})
}
